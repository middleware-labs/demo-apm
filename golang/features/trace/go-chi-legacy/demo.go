package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	mwchi "github.com/middleware-labs/agent-apm-go-chi-legacy"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
		track.WithConfigTag(track.Token, "your access token"),
	)
	// define router
	r := chi.NewRouter()
	r.Use(mwchi.Middleware("my-server", mwchi.WithChiRoutes(r)))
	r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := 1
		name := "ron"
		reply := fmt.Sprintf("user %s (id %s)\n", name, id)
		w.Write(([]byte)(reply))
	}))
	// serve router
	_ = http.ListenAndServe(":8080", r)
}
