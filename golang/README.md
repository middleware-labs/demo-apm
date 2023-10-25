# Golang APM Demo

[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/middleware-labs/golang-apm)](https://github.com/middleware-labs/golang-apm)

|  Traces  |  Metrics  |  Profiling  |  Logs (App/Custom)  |
|:--------:|:---------:|:-----------:|:-------------------:|
|   Yes    |    Yes    |     Yes     |       No/Yes        |

## Prequisites

If you are expecting golang demo data on your dashboard, make sure you have our Host Agent installed.

---------------------

## Log Collection
```
go run features/log/log.go
```
## Distributed Tracing

gin-gonic/gin demo
```
go run features/trace/gin/gin.go
```

gorilla/mux demo
```
go run features/trace/mux/mux.go
```

database/sql demo
```
go run features/trace/sql/sql.go
```

pg demo
```
go run features/trace/pg/pg.go
```

## Golang Specific Metrics
```
go run features/metric/metric.go
```

## Complete Example
```
go run main.go
```


