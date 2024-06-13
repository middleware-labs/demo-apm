## Prerequisites

To monitor APM data on a dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.17 or higher is required


## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-http
```

## Usage

```
import (
      "net/http"
      
	  track "github.com/middleware-labs/golang-apm/tracker"
	  mwhttp "github.com/middleware-labs/golang-apm-http/http"
)
```

## Configure Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with net/http.
```   
go track.Track(
    track.WithConfigTag(track.Service, "your service name"),
    track.WithConfigTag(track.Project, "your project name"),
    track.WithConfigTag(track.Token, "your access token"),
)
http.Handle("/hello", mwhttp.MiddlewareHandler(http.HandlerFunc(hello), "hello"))
```
