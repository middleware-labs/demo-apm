# Golang APM Demo

[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/middleware-labs/golang-apm)](https://github.com/middleware-labs/golang-apm)

|  Traces  |  Metrics  |  Profiling  |  Application Logs   | Custom Logs | AWS Lambda  | 
|:--------:|:---------:|:-----------:|:-------------------:|:-----------:|:-----------:|
|   Yes    |    Yes    |     Yes     |          Yes        |     Yes     |     Yes     |

## Prequisites

If you are expecting golang demo data on your dashboard, make sure you have our Host Agent installed.

---------------------

## Log Collection
```go
go run features/log/log.go
```
## Distributed Tracing

gin-gonic/gin demo
```go
go run features/trace/gin/gin.go
```

gorilla/mux demo
```go
go run features/trace/mux/mux.go
```

database/sql demo
```go
go run features/trace/sql/sql.go
```

pg demo
```go
go run features/trace/pg/pg.go
```

## Golang Specific Metrics
```go
go run features/metric/metric.go
```

## Golang Application Logs

### `log/slog`
```go
go run features/mwotelslog.go
```
### `zap`
```go
go run features/mwotelzap.go
```
### `zerolog`
```go
go run features/mwotelzerolog.go
```

## Complete Example
```go
go run main.go
```


