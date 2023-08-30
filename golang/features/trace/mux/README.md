# Prerequisites
To monitor APM data on dashboard, Middleware Host agent needs to be installed.


## Golang Version Support
go v1.17 or higher is required


## Gin Version Support 

The Latest APM version: 
`https://github.com/middleware-labs/golang-apm-mux/blob/v0.4.1`
works well with `github.com/gorilla/mux v1.8.0`


Refer https://github.com/middleware-labs/golang-apm-gin/blob/v0.4.2/go.mod


## Installation

```
go get github.com/middleware-labs/golang-apm
go get github.com/middleware-labs/golang-apm-mux
```

## Usage

```
import (
	  track "github.com/middleware-labs/golang-apm/tracker"
	  mwmux "github.com/middleware-labs/golang-apm-mux/mux"
)
```


## Configure MUX to use Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with MUX framework.
This code should be put wherever MUX package is initialized (typically in the main function of your code).

```
r := mux.NewRouter()
config, _ := track.Track(
    track.WithConfigTag("service", "your service name"),
    track.WithConfigTag("projectName", "your project name"),
    track.WithConfigTag("accessToken", "your access token"),
)
r.Use(mwmux.Middleware(config))
```