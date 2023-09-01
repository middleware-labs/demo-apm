# Prerequisites

To monitor APM data on a dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.17 or higher is required

## astaxie/beego Version Support

The Latest APM version:
`https://github.com/middleware-labs/golang-apm-beego/blob/v0.0.7`
works well with `github.com/astaxie/beego v1.12.3` and higher

## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-beego
```

## Usage

```
import (
	mw_beego "github.com/middleware-labs/golang-apm-beego/beego"
	track "github.com/middleware-labs/golang-apm/tracker"
)
```

## Configure Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with astaxie/beego framework.
```
config, _ := track.Track(__
	track.WithConfigTag("service", "your service name"),
	track.WithConfigTag("projectName", "your project name"),
	track.WithConfigTag("accessToken", "your access token"),
)
mware := mw_beego.Middleware(config)
```


