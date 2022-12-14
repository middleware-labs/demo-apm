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
---------------------

## Collect Golang specific metrics

Call track method in your main function
```
go track.Track(
    track.WithConfigTag("service", "your-service-name"),
    track.WithConfigTag("projectName", "your-project-name"),
)
```
Running this method with go routine is important !

This will start collecting the application Metrics

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
|go-pg/mongo        |   [Mongo Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/mongo)  |
|gRPC               |   [gRPC Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/grpc)  |

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