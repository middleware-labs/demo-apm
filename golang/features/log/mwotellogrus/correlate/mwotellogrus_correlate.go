package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	mwhttp "github.com/middleware-labs/golang-apm-http/http"
	otellog "github.com/middleware-labs/golang-apm/mwotellogrus"
	track "github.com/middleware-labs/golang-apm/tracker"
	log "github.com/sirupsen/logrus"
)

func StandardFields(ctx context.Context) log.Fields {
	return log.Fields{
		otellog.MWTraceID: track.TraceID(ctx),
		otellog.MWSpanID:  track.SpanID(ctx),
	}
}

func main() {

	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
	)

	logHook := otellog.NewMWOTelHook(config, otellog.WithLevels(log.AllLevels), otellog.WithName("otellogs"))

	// add hook in logrus
	log.AddHook(logHook)
	// set formatter
	log.SetFormatter(&log.JSONFormatter{})

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
	//set context in logrus for correlation
	logger := log.WithFields(StandardFields(r.Context()))

	logger.Info("Info helloHandler testing 123456")
	logger.Debug("Debug helloHandler testing 123456")
	logger.Error("Error helloHandler testing 123456")

	fmt.Fprintf(w, "Hello, World!")
}

func makeRequest() {
	for {
		time.Sleep(1 * time.Second)
		http.Get("http://localhost:8090/hello")
	}
}
