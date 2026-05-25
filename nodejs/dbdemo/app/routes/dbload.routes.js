// Routes that drive the synthetic Postgres workload (see ../dbload.js).
// Each handler runs inside the Express request span, so the DB spans it
// produces are children of the request trace and carry a traceparent SQL
// comment that shows up in pg_stat_activity.
module.exports = app => {
    const dbload = require("../dbload.js");
    const router = require("express").Router();

    const run = (fn) => (req, res) => {
        fn()
            .then((r) => {
                // Logged inside the active request span -> with consoleLog:true the
                // agent stamps this line with the same trace_id/span_id (log<->trace).
                console.log(`[dbload] ${r && r.scenario}: ${JSON.stringify(r)}`);
                res.status(200).send(r);
            })
            .catch((e) => {
                console.error(`[dbload] error: ${String((e && e.message) || e)}`);
                res.status(500).send({ error: String((e && e.message) || e) });
            });
    };

    router.get("/complex",  run(dbload.complexAnalytics));
    router.get("/blocking", run(dbload.lockingBlocking));
    router.get("/iowait",   run(dbload.ioWaitHeavySort));
    router.get("/seqscan",  run(dbload.seqScanHeavy));
    router.get("/idletx",   run(dbload.idleInTransaction));
    router.get("/deadlock", run(dbload.deadlock));
    router.get("/insert",   run(dbload.insertBatch));
    router.get("/delete",   run(dbload.deleteRandom));
    router.get("/cleanup",  run(dbload.cleanup));
    router.get("/random",   run(dbload.runRandom));

    app.use('/api/dbload', router);
};
