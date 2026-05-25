'use strict';

// Synthetic, "realistic" Postgres workload generator for the tutorials demo.
//
// Each scenario runs genuine SQL that surfaces in pg_stat_activity with a
// recognisable state (active heavy query, Lock wait, IO/temp spill, idle in
// transaction, deadlock, write churn). Because app.js enables the OTel SQL
// commenter, every statement issued here carries a /*traceparent='...'*/
// comment, so rows in pg_stat_activity link straight back to a trace.
//
// Storage safety: inserts grow the table, deletes churn it, and a background
// maintenance loop caps the row count (MAX_ROWS) and periodically VACUUMs so
// the database can never fill the disk on its own.

const { Pool } = require('pg');

// Dedicated pool with a distinct application_name so these queries are trivial
// to find:  SELECT * FROM pg_stat_activity WHERE application_name='tutorials-dbload';
const pool = new Pool({
    host: process.env.POSTGRES_HOST || 'localhost',
    port: parseInt(process.env.POSTGRES_PORT || '5432', 10),
    user: process.env.POSTGRES_USER || 'postgres',
    password: process.env.POSTGRES_PASSWORD || 'postgres',
    database: process.env.POSTGRES_DB || 'todo',
    max: 15,
    idleTimeoutMillis: 30000,
    application_name: 'tutorials-dbload',
});

pool.on('error', (err) => console.error('[dbload] idle client error:', err.message));

const rnd = (min, max) => Math.floor(Math.random() * (max - min + 1)) + min;
const sleep = (ms) => new Promise((res) => setTimeout(res, ms));

const ROW_TARGET = 100000;   // baseline number of rows seeded at startup
const MAX_ROWS   = 250000;   // hard ceiling enforced by the maintenance loop

// ---------------------------------------------------------------------------
// One-time schema + seed so scans / sorts / row locks have real data to chew on
// ---------------------------------------------------------------------------
let readyPromise = null;
function ready() {
    if (!readyPromise) {
        readyPromise = ensureSchema().catch((e) => {
            console.error('[dbload] ensureSchema failed:', e.message);
            readyPromise = null; // allow a retry on the next call
            throw e;
        });
    }
    return readyPromise;
}

async function ensureSchema() {
    await pool.query(`
        CREATE TABLE IF NOT EXISTS dbload_items (
            id          serial PRIMARY KEY,
            category    text          NOT NULL,
            amount      numeric(12,2) NOT NULL,
            status      text          NOT NULL,
            description text          NOT NULL,
            created_at  timestamptz   NOT NULL DEFAULT now()
        )`);
    const { rows } = await pool.query('SELECT count(*)::int AS c FROM dbload_items');
    const missing = ROW_TARGET - rows[0].c;
    if (missing > 0) {
        console.log(`[dbload] seeding dbload_items with ${missing} rows...`);
        await seedRows(missing);
        await pool.query('CREATE INDEX IF NOT EXISTS idx_dbload_items_cat ON dbload_items(category)');
        await pool.query('ANALYZE dbload_items');
    }
    console.log('[dbload] schema ready');
}

function seedRows(n) {
    return pool.query(`
        INSERT INTO dbload_items (category, amount, status, description)
        SELECT (ARRAY['alpha','beta','gamma','delta','omega'])[1 + (g % 5)],
               round((random() * 1000)::numeric, 2),
               (ARRAY['new','open','pending','closed','archived'])[1 + (g % 5)],
               md5(g::text) || '-' || md5(random()::text)
        FROM generate_series(1, $1) g`, [n]);
}

// ---------------------------------------------------------------------------
// Scenarios
// ---------------------------------------------------------------------------

// 1) Heavy analytical read: window functions + percentile + group/join + sort,
//    then a short pg_sleep so it lingers as an "active" query in pg_stat_activity.
async function _complexAnalytics() {
    const client = await pool.connect();
    try {
        const t = Date.now();
        const res = await client.query(`
            WITH ranked AS (
                SELECT category, status, amount,
                       row_number() OVER (PARTITION BY category ORDER BY amount DESC) AS rn,
                       avg(amount)  OVER (PARTITION BY category)                       AS cat_avg
                FROM dbload_items
            ),
            agg AS (
                SELECT category, status, count(*) AS n, max(amount) AS max_amount,
                       percentile_cont(0.95) WITHIN GROUP (ORDER BY amount) AS p95
                FROM dbload_items
                GROUP BY category, status
            )
            SELECT r.category, r.status, r.rn, r.cat_avg, a.n, a.max_amount, a.p95
            FROM ranked r
            JOIN agg a ON a.category = r.category AND a.status = r.status
            WHERE r.rn <= 5
            ORDER BY a.p95 DESC, r.category
            LIMIT 50`);
        const naps = (rnd(400, 1500) / 1000).toFixed(3);
        await client.query('SELECT pg_sleep($1)', [naps]);
        return { scenario: 'complexAnalytics', rows: res.rowCount, sleptSec: Number(naps), ms: Date.now() - t };
    } finally {
        client.release();
    }
}

// 2) Lock contention: txn A holds a row lock while txn B waits on the same row.
//    B shows up in pg_stat_activity with wait_event_type='Lock'.
async function _lockingBlocking() {
    const a = await pool.connect();
    const b = await pool.connect();
    const id = rnd(1, ROW_TARGET);
    const holdMs = rnd(2000, 5000);
    try {
        await a.query('BEGIN');
        await a.query('UPDATE dbload_items SET amount = amount + 1, status = $2 WHERE id = $1', [id, 'pending']);
        // Fire B's update on the same row WITHOUT awaiting — it blocks until A commits.
        const bUpdate = b.query('UPDATE dbload_items SET amount = amount + 2 WHERE id = $1', [id]);
        await sleep(holdMs);            // <- B is blocked for this whole window
        await a.query('COMMIT');
        await bUpdate;                  // B unblocks and autocommits
        return { scenario: 'lockingBlocking', id, blockedMs: holdMs };
    } catch (e) {
        try { await a.query('ROLLBACK'); } catch (_) {}
        throw e;
    } finally {
        a.release();
        b.release();
    }
}

// 3) IO / temp-file pressure: a tiny work_mem forces a large sort to spill to
//    disk (wait_event_type='IO', BufFileRead/Write; temp bytes reported).
//    Temp files are released automatically when the query ends.
async function _ioWaitHeavySort() {
    const client = await pool.connect();
    try {
        const t = Date.now();
        await client.query("SET work_mem = '64kB'");
        const n = rnd(800000, 1800000);
        const res = await client.query(`
            SELECT g, md5(g::text) AS h
            FROM generate_series(1, $1) g
            ORDER BY md5(g::text)
            LIMIT 100`, [n]);
        await client.query('RESET work_mem');
        return { scenario: 'ioWaitHeavySort', sorted: n, rows: res.rowCount, ms: Date.now() - t };
    } finally {
        client.release();
    }
}

// 4) CPU-bound sequential scan over the whole table (per-row md5).
async function _seqScanHeavy() {
    const client = await pool.connect();
    try {
        const t = Date.now();
        const res = await client.query(`
            SELECT count(*) AS matches
            FROM dbload_items
            WHERE md5(description) < md5(category || amount::text)
              AND description LIKE '%a%'`);
        return { scenario: 'seqScanHeavy', matches: Number(res.rows[0].matches), ms: Date.now() - t };
    } finally {
        client.release();
    }
}

// 5) Idle-in-transaction: open a txn, run a query, sit idle, then commit.
//    Shows as state='idle in transaction' in pg_stat_activity.
async function _idleInTransaction() {
    const client = await pool.connect();
    const idleMs = rnd(3000, 6000);
    try {
        await client.query('BEGIN');
        await client.query('SELECT * FROM dbload_items WHERE id = $1', [rnd(1, ROW_TARGET)]);
        await sleep(idleMs);
        await client.query('COMMIT');
        return { scenario: 'idleInTransaction', idleMs };
    } catch (e) {
        try { await client.query('ROLLBACK'); } catch (_) {}
        throw e;
    } finally {
        client.release();
    }
}

// 6) Deadlock: two txns lock two rows in opposite order; Postgres aborts one
//    with 'deadlock detected' (40P01), producing an error span. Expected, so we
//    swallow it and report which rows were involved.
async function _deadlock() {
    const a = await pool.connect();
    const b = await pool.connect();
    let id1 = rnd(1, ROW_TARGET), id2 = rnd(1, ROW_TARGET);
    while (id2 === id1) id2 = rnd(1, ROW_TARGET);
    let deadlockHit = false;
    try {
        await a.query('BEGIN');
        await b.query('BEGIN');
        await a.query('UPDATE dbload_items SET amount = amount + 1 WHERE id = $1', [id1]);
        await b.query('UPDATE dbload_items SET amount = amount + 1 WHERE id = $1', [id2]);
        // Cross-lock: each now waits for the row the other holds -> deadlock.
        const ax = a.query('UPDATE dbload_items SET amount = amount + 1 WHERE id = $1', [id2]);
        const bx = b.query('UPDATE dbload_items SET amount = amount + 1 WHERE id = $1', [id1]);
        const results = await Promise.allSettled([ax, bx]);
        deadlockHit = results.some((r) => r.status === 'rejected' && /deadlock/i.test(String(r.reason)));
        try { await a.query('COMMIT'); } catch (_) {}
        try { await b.query('COMMIT'); } catch (_) {}
        return { scenario: 'deadlock', id1, id2, deadlockDetected: deadlockHit };
    } finally {
        try { await a.query('ROLLBACK'); } catch (_) {}
        try { await b.query('ROLLBACK'); } catch (_) {}
        a.release();
        b.release();
    }
}

// 7) Write churn — INSERT: append a batch of fresh rows.
async function _insertBatch() {
    const n = rnd(500, 2500);
    const t = Date.now();
    const res = await seedRows(n);
    return { scenario: 'insertBatch', inserted: res.rowCount, ms: Date.now() - t };
}

// 8) Write churn — DELETE: remove a random batch of rows (full scan + sort).
async function _deleteRandom() {
    const client = await pool.connect();
    const n = rnd(200, 1500);
    try {
        const t = Date.now();
        const res = await client.query(`
            DELETE FROM dbload_items
            WHERE id IN (SELECT id FROM dbload_items ORDER BY random() LIMIT $1)`, [n]);
        return { scenario: 'deleteRandom', deleted: res.rowCount, ms: Date.now() - t };
    } finally {
        client.release();
    }
}

// ---------------------------------------------------------------------------
// Automatic storage maintenance: keep the table under MAX_ROWS and reclaim
// dead-tuple bloat with periodic VACUUM so the DB can't fill the disk.
// ---------------------------------------------------------------------------
async function enforceCap() {
    await ready();
    const { rows } = await pool.query('SELECT count(*)::int AS c FROM dbload_items');
    const excess = rows[0].c - MAX_ROWS;
    let deleted = 0;
    if (excess > 0) {
        const res = await pool.query(
            `DELETE FROM dbload_items
             WHERE id IN (SELECT id FROM dbload_items ORDER BY id ASC LIMIT $1)`, [excess]);
        deleted = res.rowCount;
        console.log(`[dbload] cap cleanup: deleted ${deleted} rows (was ${rows[0].c}, cap ${MAX_ROWS})`);
    }
    return { scenario: 'cleanup', rowsBefore: rows[0].c, deleted, cap: MAX_ROWS };
}

let maintenanceTicks = 0;
async function maintenance() {
    try {
        await enforceCap();
        maintenanceTicks++;
        // VACUUM (ANALYZE) every ~5 minutes to reclaim dead tuples for reuse.
        if (maintenanceTicks % 7 === 0) {
            await pool.query('VACUUM (ANALYZE) dbload_items');
            console.log('[dbload] VACUUM (ANALYZE) dbload_items done');
        }
    } catch (e) {
        console.error('[dbload] maintenance error:', e.message);
    }
}

// ---------------------------------------------------------------------------
// Concurrency guard + weighted random picker
// ---------------------------------------------------------------------------
const MAX_INFLIGHT = 8;
let inFlight = 0;

function guarded(name, fn) {
    return async function () {
        if (inFlight >= MAX_INFLIGHT) {
            return { scenario: name, skipped: true, reason: 'max in-flight reached', inFlight };
        }
        inFlight++;
        try {
            await ready();
            return await fn();
        } finally {
            inFlight--;
        }
    };
}

const complexAnalytics  = guarded('complexAnalytics',  _complexAnalytics);
const lockingBlocking   = guarded('lockingBlocking',   _lockingBlocking);
const ioWaitHeavySort   = guarded('ioWaitHeavySort',   _ioWaitHeavySort);
const seqScanHeavy      = guarded('seqScanHeavy',      _seqScanHeavy);
const idleInTransaction = guarded('idleInTransaction', _idleInTransaction);
const deadlock          = guarded('deadlock',          _deadlock);
const insertBatch       = guarded('insertBatch',       _insertBatch);
const deleteRandom      = guarded('deleteRandom',      _deleteRandom);
const cleanup           = guarded('cleanup',           enforceCap);

const SCENARIOS = [
    { fn: complexAnalytics,  w: 22 },
    { fn: lockingBlocking,   w: 16 },
    { fn: ioWaitHeavySort,   w: 14 },
    { fn: seqScanHeavy,      w: 12 },
    { fn: idleInTransaction, w: 10 },
    { fn: insertBatch,       w: 14 },  // inserts > deletes so the table stays populated;
    { fn: deleteRandom,      w: 8  },  // the maintenance loop enforces the MAX_ROWS ceiling
    { fn: deadlock,          w: 4  },
];

async function runRandom() {
    const total = SCENARIOS.reduce((s, x) => s + x.w, 0);
    let pick = Math.random() * total;
    for (const s of SCENARIOS) {
        pick -= s.w;
        if (pick <= 0) return s.fn();
    }
    return SCENARIOS[0].fn();
}

// Build the schema in the background at startup (don't block module load) and
// start the automatic storage-maintenance loop (every 45s).
ready().catch(() => {});
setInterval(maintenance, 45000).unref();

module.exports = {
    complexAnalytics, lockingBlocking, ioWaitHeavySort, seqScanHeavy,
    idleInTransaction, deadlock, insertBatch, deleteRandom, cleanup, runRandom,
};
