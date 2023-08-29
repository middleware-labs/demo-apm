# Prerequisites

- To monitor APM data on dashboard, Middleware Host agent needs to be installed.


## Install Golang package

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-http
```

## Import Middleware Tracker & HTTP Golang package

```
import (
	  track "github.com/middleware-labs/golang-apm/tracker"
	  mwhttp "github.com/middleware-labs/golang-apm-http/http"
)
```

## Add this snippet in your main function
```   
go track.Track(
    track.WithConfigTag("service", "your service name"),
    track.WithConfigTag("projectName", "your project name"),
    track.WithConfigTag("accessToken", "your access token"),
)
http.Handle("/hello", mwhttp.MiddlewareHandler(http.HandlerFunc(hello), "hello"))
```
