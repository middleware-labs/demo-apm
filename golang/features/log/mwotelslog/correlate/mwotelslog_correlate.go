package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	mwhttp "github.com/middleware-labs/golang-apm-http/http"
	"github.com/middleware-labs/golang-apm/mwotelslog"
	track "github.com/middleware-labs/golang-apm/tracker"
)

var Logger *slog.Logger

func StandardFields(ctx context.Context) []slog.Attr {
	fields := []slog.Attr{
		slog.String(mwotelslog.MWTraceID, track.TraceID(ctx)),
		slog.String(mwotelslog.MWSpanID, track.SpanID(ctx)),
	}
	return fields
}

func main() {

	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
	)

	Logger = mwotelslog.NewMWOTelLogger(
		config,
		mwotelslog.WithDefaultConsoleLog(), // to enable console log
		mwotelslog.WithName("otelslog"),
	)

	//configure default logger

	slog.SetDefault(Logger)

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

	//set context in slog for correlation
	attrs := StandardFields(r.Context())
	rlog := Logger.With(attrs[0], attrs[1])

	// start logs
	rlog.Info("Info helloHandler testing 123456")
	rlog.Debug("Debug helloHandler testing 123456")
	rlog.Error("Error helloHandler testing 123456")
	fmt.Fprintf(w, "Hello, World!")
}

func makeRequest() {
	for {
		time.Sleep(1 * time.Second)
		http.Get("http://localhost:8090/hello")
	}
}
