## Prerequisites

To monitor APM data on a dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.17 or higher is required

## go-pg Version Support 

The Latest APM version: 
`https://github.com/middleware-labs/golang-apm-pg/blob/v0.0.1`
works well with `github.com/go-pg/pg/v10 v10.11.1` and higher


## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-pg
```

## Usage

```
import (
      "github.com/go-pg/pg/v10"
	  track "github.com/middleware-labs/golang-apm/tracker"
	  mw_pg "github.com/middleware-labs/golang-apm-pg/pg"
)
```

## Configure Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with go-pg/pg/v10.

```
config, _ := track.Track(
    track.WithConfigTag(track.Service, "your service name"),
    track.WithConfigTag(track.Project, "your project name"),
    track.WithConfigTag(track.Token, "your access token"),
)
db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "postgres",
		Database: "dbname",
	})
mw_pg.AddQueryHook(db)
```
