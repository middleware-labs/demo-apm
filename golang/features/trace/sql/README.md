## Prerequisites

To monitor APM data on a dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.17 or higher is required

## SQL Version Support 

The Latest APM version: 
`https://github.com/middleware-labs/golang-apm-sql/blob/v0.0.10`
works well with `github.com/go-sql-driver/mysql v1.7.1` and higher


## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-sql
```

## Usage

```
import (
	  track "github.com/middleware-labs/golang-apm/tracker"
	  mw_sql "github.com/middleware-labs/golang-apm-sql/sql"
)
```

## Configure Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with SQL driver.

```
config, _ := track.Track(
    track.WithConfigTag("service", "your service name"),
    track.WithConfigTag("projectName", "your project name"),
    track.WithConfigTag("accessToken", "your access token"),
)
db, err := mw_sql.Open("mysql", "uname:password@tcp(127.0.0.1:3306)/todo")
```
