package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	mwchi "github.com/middleware-labs/golang-apm-go-chi"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	track.Track(
		track.WithConfigTag("service", "your service name"),
		track.WithConfigTag("projectName", "your project name"),
		track.WithConfigTag("accessToken", "Your access token"),
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
