# Prerequisites

To monitor APM data on dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.17 or higher is required

## Gin Version Support 

The Latest APM version: 
`https://github.com/middleware-labs/golang-apm-gin/blob/v0.4.2`
works well with `github.com/gin-gonic/gin v1.8.1` and higher

Refer https://github.com/middleware-labs/golang-apm-gin/blob/v0.4.2/go.mod


## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-gin
```

## Usage

```
import (
	  track "github.com/middleware-labs/golang-apm/tracker"
	  mwgin "github.com/middleware-labs/golang-apm-gin/gin"
)
```

## Configure GIN to use Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with GIN framework.
This code should be put wherever GIN package is initialized (typically in the main function of your code).

```
r := gin.Default()
config, _ := track.Track(
    track.WithConfigTag(track.Service, "your service name"),
    track.WithConfigTag(track.Project, "your project name"),
    track.WithConfigTag(track.Token, "your access token"),
)
r.Use(mwgin.Middleware(config))
```