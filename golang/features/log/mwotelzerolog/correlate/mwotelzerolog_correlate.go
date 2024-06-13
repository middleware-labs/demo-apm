package main

import (
	"fmt"
	"net/http"
	"time"

	mwhttp "github.com/middleware-labs/golang-apm-http/http"
	"github.com/middleware-labs/golang-apm/mwotelzerolog"
	track "github.com/middleware-labs/golang-apm/tracker"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger zerolog.Logger

func main() {

	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
	)

	logHook := mwotelzerolog.NewMWOTelHook(config)

	// add hook in zerolog
	logger = log.Hook(logHook)

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

	//set context in zerolog for correlation
	logger.Info().
		Str(mwotelzerolog.MWTraceID, track.TraceID(r.Context())).
		Str(mwotelzerolog.MWSpanID, track.SpanID(r.Context())).
		Msg("Info helloHandler testing 123456")
	logger.Debug().
		Str(mwotelzerolog.MWTraceID, track.TraceID(r.Context())).
		Str(mwotelzerolog.MWSpanID, track.SpanID(r.Context())).
		Msg("Debug helloHandler testing 123456")
	logger.Error().
		Str(mwotelzerolog.MWTraceID, track.TraceID(r.Context())).
		Str(mwotelzerolog.MWSpanID, track.SpanID(r.Context())).
		Msg("Error helloHandler testing 123456")

	fmt.Fprintf(w, "Hello, World!")
}

func makeRequest() {
	for {
		time.Sleep(1 * time.Second)
		http.Get("http://localhost:8090/hello")
	}
}
