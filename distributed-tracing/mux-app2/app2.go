package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"

	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	config, _ := track.Track(
		track.WithConfigTag("service", "go-app-b"),
		track.WithConfigTag("projectName", "go-app-b"),
	)
	tp := config.Tp

	r := mux.NewRouter()
	r.Use(otelmux.Middleware("mux-server-b", otelmux.WithTracerProvider(tp)))
	r.HandleFunc("/{id:[0-9]+}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("ID:", id)
		_, _ = w.Write(([]byte)(id))
	}))
	http.Handle("/", r)
	_ = http.ListenAndServe(":3001", nil) //nolint:gosec // Ignoring G114: Use of net/http serve function that has no support for setting timeouts.
}
