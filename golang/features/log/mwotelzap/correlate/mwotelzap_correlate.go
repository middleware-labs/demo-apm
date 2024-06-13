package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	mwhttp "github.com/middleware-labs/golang-apm-http/http"
	mwotelzap "github.com/middleware-labs/golang-apm/mwotelzap"
	track "github.com/middleware-labs/golang-apm/tracker"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func StandardFields(ctx context.Context) []zapcore.Field {
	fields := []zapcore.Field{
		zapcore.Field{Key: mwotelzap.MWTraceID, String: track.TraceID(ctx), Type: zapcore.StringType},
		zapcore.Field{Key: mwotelzap.MWSpanID, String: track.SpanID(ctx), Type: zapcore.StringType},
	}
	return fields
}

func main() {

	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
	)

	// Define encoder configuration
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Configure console output
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)
	consoleLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleDebugging, consoleLevel)

	// Configure file output
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}
	defer file.Close()
	fileEncoder := zapcore.NewConsoleEncoder(encoderCfg)  // You can use a different encoder if needed
	fileLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel) // Set your desired log level here
	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(file), fileLevel)

	logger = zap.New(zapcore.NewTee(consoleCore, fileCore, mwotelzap.NewMWOTelCore(config, mwotelzap.WithName("otelzaplog"))))
	zap.ReplaceGlobals(logger)

	//use mwhttp for http handler instrumentation
	http.Handle("/hello", mwhttp.MiddlewareHandler(http.HandlerFunc(helloHandler), "hello"))
	fmt.Println("listening on 8090")

	// this make continuos requests
	go makeRequest()

	// start the server
	http.ListenAndServe(":8090", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	time.Sleep(2 * time.Second)

	// set context in zap for correlation
	zlog := logger.With(StandardFields(r.Context())...)

	zlog.Error("Error zap")
	zlog.Info("Info zap")
	zlog.Warn("Warn zap")
}

func makeRequest() {
	for {
		time.Sleep(3 * time.Second)
		http.Get("http://localhost:8090/hello")
	}
}
