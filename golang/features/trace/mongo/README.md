## Prerequisites

To monitor APM data on a dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.17 or higher is required

## Mongo Version Support 

The Latest APM version: 
`https://github.com/middleware-labs/golang-apm-mongo/blob/v0.0.3`
works well with `go.mongodb.org/mongo-driver v1.12.1` and higher


## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-mongo
```

## Usage

```
import (
	 "go.mongodb.org/mongo-driver/mongo/options"
	  
	  track "github.com/middleware-labs/golang-apm/tracker"
	  mw_mongo "github.com/middleware-labs/golang-apm-mongo/mongo"
)
```

## Configure Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with Mongo driver.

```
config, _ := track.Track(
    track.WithConfigTag(track.Service, "your service name"),
    track.WithConfigTag(track.Project, "your project name"),
    track.WithConfigTag(track.Token, "your access token"),
)
opts := options.Client()
opts.Monitor = mw_mongo.NewMonitor()
```
