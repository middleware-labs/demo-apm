## Prerequisites

To monitor APM data on a dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.17 or higher is required.

# GORM 2.0 Version Support

The Latest APM version:
`https://github.com/middleware-labs/go-agent-gorm/blob/v0.0.6`
works well with `gorm.io/gorm v1.25.1` and higher


## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/go-agent-gorm
```

## Usage

```
import (
      "gorm.io/gorm"
      "gorm.io/driver/mysql"
	  track "github.com/middleware-labs/golang-apm/tracker"
	  mw_gorm "github.com/middleware-labs/go-agent-gorm/gorm"
)
```

## Configure Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with GORM 2.0.

```
config, _ := track.Track(
    track.WithConfigTag("service", "your service name"),
    track.WithConfigTag("projectName", "your project name"),
    track.WithConfigTag("accessToken", "your access token"),
)
db, err := gorm.Open(mysql.Open("username:password@tcp(127.0.0.1:3306)/dbname"), &gorm.Config{})
if err != nil {
   panic(err)
}
if err := db.Use(mw_gorm.NewPlugin(mw_gorm.WithDBName("dbname"), mw_gorm.WithAttributes(track.String("db.system", "mysql")))); err != nil {
  panic(err)
}
```

And then use db.WithContext(ctx) to propagate the active span

```
var num int
if err := db.WithContext(ctx).Raw("SELECT 42").Scan(&num).Error; err != nil {
	panic(err)
}
```
