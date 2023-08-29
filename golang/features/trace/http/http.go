package main

import (
	"fmt"
	"net/http"

	mwhttp "github.com/middleware-labs/golang-apm-http/http"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	go track.Track(
		track.WithConfigTag("service", "your service name"),
		track.WithConfigTag("projectName", "your project name"),
		track.WithConfigTag("accessToken", "your access token"),
	)
	http.Handle("/hello", mwhttp.MiddlewareHandler(http.HandlerFunc(hello), "hello"))
	fmt.Println("listening on 8090")
	http.ListenAndServe(":8090", nil)
}
