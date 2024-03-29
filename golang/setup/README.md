# Golang APM Setup

## Prerequisites

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

## Collect Golang traces

Call track method in your main function
```
go track.Track(
    track.WithConfigTag("service", {APM-SERVICE-NAME}),
    track.WithConfigTag("projectName", {APM-PROJECT-NAME}),
)
```
Running this method with go routine is important !

This will start collecting the application traces

## Collect Application Profiling Data

If you also want to collect profiling data for your application,
simply add this one config to your track.Track() call

```
track.WithConfigTag("accessToken", "{ACCOUNT_KEY}")
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
|go-chi             |   [go-chi Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/go-chi-legacy)  |
|go-chi/v5          |   [go-chi/v5 Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/go-chi)  |
|astaxie-beego      |   [astaxie-beego Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/astaxie-beego)  |
|beego/v2           |   [beego/v2 Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/beego-v2)  |
|database/sql       |   [SQL Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/sql)  |
|go-pg/pg           |   [PG Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/pg)  |
|GORM 1             |   [GORM 1 Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/gorm1)  |
|GORM 2             |   [GORM 2 Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/gorm2)  |
|go-pg/mongo        |   [Mongo Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/mongo)  |
|gRPC               |   [gRPC Demo](https://github.com/middleware-labs/demo-apm/tree/master/golang/features/trace/grpc)  |

---------------

## Note for APM inside Kubernetes

If you are using APM in a Kubernetes cluster make sure to follow these 2 steps:

### Step 1 : Find your Middleware Service namespace
For older setup, your "mw-service" can be inside "mw-agent-ns-{FIRST-5-LETTERS-OF-API-KEY}" namespace

For newer setup, we simplified the namespace name to "mw-agent-ns"

### Step 2 : Set this ENV variable in your application deployment YAML
```
MW_AGENT_SERVICE=mw-service.NAMESPACE.svc.cluster.local
```
Please replace "NAMESPACE" with the correct value that you found from Step 1.

----------------

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
