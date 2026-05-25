const tracker = require('@middleware.io/node-apm');

// Start the Middleware agent.
//  - disabledInstrumentations:"pg"  -> turn off the agent's built-in pg
//    auto-instrumentation so we can register our own with the SQL commenter
//    (the agent exposes no knob for addSqlCommenterCommentToQueries).
//  - consoleLog:true                -> route console.log through the agent's
//    OTel log pipeline, which stamps each record with the active span's
//    trace_id/span_id (log <-> trace correlation).
//
// Note on the other two correlation needs: both are already automatic with the
// Middleware agent, so no manual code here:
//   * cross-service propagation -> the agent's global propagator + HTTP
//     auto-instrumentation inject/extract trace context on every request.
//   * logs <-> traces           -> the agent stamps trace_id/span_id on logs
//     emitted within a span (tracker.info/error, console.error, winston/pino).
//     consoleLog:true above just routes plain console.log through that pipeline.
tracker.track({ disabledInstrumentations: "pg", consoleLog: true });

// ---- Correlation: DB query <-> trace via SQL commenter -------------------
// Register pg instrumentation with the W3C SQL commenter + enhanced reporting.
// This patches `pg` (via require-in-the-middle) BEFORE the DB layer requires it
// further down, so every statement is suffixed with a comment like
//   /*traceparent='00-<trace-id>-<span-id>-01'*/
// which shows up in pg_stat_activity.query and links the running query back to
// its trace. Spans still flow to Middleware through the global TracerProvider
// that the agent's NodeSDK installed in track() above.
const { registerInstrumentations } = require('@opentelemetry/instrumentation');
const { PgInstrumentation } = require('@opentelemetry/instrumentation-pg');
registerInstrumentations({
    instrumentations: [
        new PgInstrumentation({
            addSqlCommenterCommentToQueries: true,
            enhancedDatabaseReporting: true,
            requireParentSpan: false,
        }),
    ],
});

const express = require('express');
const app = express()
const port = 3002
const cors = require("cors");

app.use(cors({origin: `http://localhost:${port}`}));

app.use(express.json()); /* bodyParser.json() is deprecated */

app.use(express.urlencoded({ extended: true })); /* bodyParser.urlencoded() is deprecated */

app.get('/304-error', async function (req, res) {
    try{
        throw new Error('Not Modified!');
    }catch (e) {
        tracker.errorRecord(e)
    }
    return res.status(304).send('Not Modified');
})

app.get('/500-error', async function (req, res) {
    try{
        throw new Error('Internal Server Occurred');
    }catch (e) {
        tracker.errorRecord(e)
    }
    return res.status(500).send('Internal Server Occurred');
})

app.get('/504-error', async function (req, res) {
    try{
        throw new Error('Timeout error');
    }catch (e) {
        tracker.errorRecord(e)
    }
    return res.status(504).send('Timeout error');
})

app.listen(port, () => {
    console.log(`Listening movies at http://localhost:${port}`)
})

if (process.env.MW_AUTOGENERATE_TRACING_DATA) {
    require("./tracingloop")
}

require("./app/routes/tutorial.routes.js")(app);
require("./app/routes/dbload.routes.js")(app);
