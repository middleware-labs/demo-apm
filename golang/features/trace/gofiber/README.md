## Prerequisites

To monitor APM data on a dashboard, Middleware Host agent needs to be installed.

## Golang Version Support

go v1.19 or higher is required


## Installation

```
go get github.com/middleware-labs/golang-apm
go get -u github.com/gofiber/contrib/otelfiber/v2
```

## Usage

```
import (
	  track "github.com/middleware-labs/golang-apm/tracker"
	  github.com/gofiber/contrib/otelfiber
      "github.com/gofiber/fiber/v2"
)
```

## Configure Middleware APM

The code snippet below initializes Middleware Go APM package and shows how to use it with gofiber/fiber/v2.

```
config, _ := track.Track(
    track.WithConfigTag(track.Service, "your service name"),
    track.WithConfigTag(track.Token, "your access token"),
)
app := fiber.New()
app.Use(otelfiber.Middleware())

```
