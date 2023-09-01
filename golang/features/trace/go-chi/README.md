# Prerequisites

To monitor APM data on a dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.17 or higher is required

## go-chi/chi/v5 Version Support 

The Latest APM version: 
`https://github.com/middleware-labs/golang-apm-go-chi/blob/v0.0.2`
works well with `github.com/go-chi/chi/v5 v5.0.8` and higher


## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-go-chi
```

## Usage

```
import (
	  track "github.com/middleware-labs/golang-apm/tracker"
	  mwchi "github.com/middleware-labs/golang-apm-go-chi"
)
```

## Configure Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with go-chi/chi/v5.

```
config, _ := track.Track(
    track.WithConfigTag("service", "your service name"),
    track.WithConfigTag("projectName", "your project name"),
    track.WithConfigTag("accessToken", "your access token"),
)
r := chi.NewRouter()
r.Use(mwchi.Middleware("my-server", mwchi.WithChiRoutes(r)))
```
