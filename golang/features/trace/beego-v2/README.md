# Prerequisites

To monitor APM data on a dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.17 or higher is required

## beego/beego/v2 Version Support

The Latest APM version:
`https://github.com/middleware-labs/golang-apm-beego-beego/blob/v0.0.4`
works well with `github.com/beego/beego/v2 v2.0.7` and higher


## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-beego-beego
```

## Usage

```
import (
	mwbeego "github.com/middleware-labs/golang-apm-beego-beego"
	track "github.com/middleware-labs/golang-apm/tracker"
)
```

## Configure Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with beego/beego/v2 framework.

```
config, _ := track.Track(
   track.WithConfigTag("service", "your service name"),
   track.WithConfigTag("projectName", "your project name"),
   track.WithConfigTag("accessToken", "your access token"),
)
mware := mwbeego.MiddleWare(config.ServiceName)
```


