## Prerequisites

To monitor APM data on a dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.17 or higher is required.

# GRPC Version Support

The Latest APM version:
`https://github.com/middleware-labs/golang-apm-grpc/blob/v0.0.3`
works well with `google.golang.org/grpc v1.53.0` and higher


## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-grpc
```

## Usage

```
import (
      track "github.com/middleware-labs/golang-apm/tracker"
	  mw_grpc "github.com/middleware-labs/golang-apm-grpc/grpc
)
```

## Configure Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with GRPC.

```
config, _ := track.Track(
    track.WithConfigTag(track.Service, "your service name"),
    track.WithConfigTag(track.Project, "your project name"),
    track.WithConfigTag(track.Token, "your access token"),
)
server := grpc.NewServer(
	grpc.UnaryInterceptor(mw_grpc.UnaryServerInterceptor()),
	grpc.StreamInterceptor(mw_grpc.StreamServerInterceptor())
)
```
