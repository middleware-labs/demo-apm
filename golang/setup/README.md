# Golang APM Setup

## Prequisites :

* To monitor APM data on dashboard, Middleware Host agent needs to be installed.
* You can refer [this demo project](https://github.com/middleware-labs/demo-apm/tree/master/golang) to refer use cases of APM.

--------------------

## Step 1 : Install Golang package

Run this in your terminal
```
go get github.com/middleware-labs/golang-apm
```

## Step 2 : Import Tracker

Add these line at the very start of your project

```
import (
    track "github.com/middleware-labs/golang-apm/tracker"
)
```
Setup "Project Name" & "Service Name"
```
go track.Track(
    track.WithConfigTag("service", "your-service-name"),
    track.WithConfigTag("projectName", "your-project-name"),
)
```
---------------------

## Collect Application specific metrics

By default we record some basic metrics for Golang. Follow the steps given below to add custom metrics.

### Step 1: Create MW meter

Add this snippet
```
meter := track.Meter()
``` 

### Step 2: Select Relevant Instrument & Start Recording Data

We support all the OTEL specified metric types with our custom meter
https://pkg.go.dev/go.opentelemetry.io/otel/metric@v0.31.0#Meter

This includes ....

|    Type    |   Instruments   |
| ------ | ---- |
| Synchronous   | Counter, UpDownCounter, Histogram |
| Asynchronous  | Counter, UpDownCounter, Gauge |

Data collected through MW meter can then be  visualized in Middleware UI

Example Snippets are as follows ... 

```
counter, err := meter.SyncFloat64().Counter("foo")
if err != nil {
    log.Fatalf("Failed to create the foo instrument: %v", err)
}
counter.Add(ctx, 1.45)
```
```
histogram, err := meter.SyncFloat64().Histogram("baz")
	if err != nil {
		log.Fatal(err)
	}
histogram.Record(ctx, 23)
```
```
gauge, err := meter.AsyncFloat64().Gauge("bar")
	if err != nil {
		log.Fatal(err)
	}
gauge.Observe(ctx,45.5)
```



## Add custom logs

```
"github.com/middleware-labs/golang-apm/logger"

....

logger.Error("Error")
logger.Info("Info")
logger.Warn("Warn")
```

## Distributed Tracing

You may need to add a framework specific middleware, to watch traces.

|Framework  |   Reference   |
|------             |    ---------  |
|gin/gonic          |   [GIN Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/gin)   |
|gorilla/mux        |   [MUX Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/mux)  |
|database/sql       |   [SQL Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/sql)  |
|go-pg/pg           |   [PG Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/pg)  |

---------------

## Note :

If you are using APM in a Kubernetes cluster, Make sure to pass this ENV variable:

```
MW_AGENT_SERVICE=mw-service.mw-agent-ns-{FIRST-5-LETTERS-OF-API-KEY}.svc.cluster.local
```

## Error Handling :

If you want to record exception in traces then you can use track.RecordError(ctx,error) method.

```golang

app.get('/hello', (req, res) => {
    ctx := req.Context()
    try {
        throw ("error");
    } catch (error) {
        track.RecordError(ctx, error)
    }
})

```