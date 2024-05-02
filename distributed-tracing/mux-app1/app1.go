package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	track "github.com/middleware-labs/golang-apm/tracker"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("mux-server-1")

func main() {
	config, _ := track.Track(
		track.WithConfigTag("service", "go-app-a"),
		track.WithConfigTag("projectName", "go-app-a"),
		track.WithConfigTag("accessToken", "nicugewzpmoxkjrobuclnywdfuerxwavunka"),
	)
	tp := config.Tp
	r := mux.NewRouter()
	r.Use(otelmux.Middleware("mux-server-a", otelmux.WithTracerProvider(tp)))

	r.HandleFunc("/users/{id:[0-9]+}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		name := getUser(r.Context(), id)
		reply := fmt.Sprintf("user %s (id %s)\n", name, id)
		_, _ = w.Write(([]byte)(reply))
	}))
	http.Handle("/", r)
	_ = http.ListenAndServe(":8090", nil) //nolint:gosec // Ignoring G114: Use of net/http serve function that has no support for setting timeouts.
}

func getUser(ctx context.Context, id string) string {
	_, span := tracer.Start(ctx, "getUser", oteltrace.WithAttributes(attribute.String("id", id)))
	defer span.End()
	request, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:3001/"+id, nil)

	client := http.Client{
		// Wrap the Transport with one that starts a span and injects the span context
		// into the outbound request headers.
		Transport: otelhttp.NewTransport(http.DefaultTransport),
		Timeout:   10 * time.Second,
	}
	resp, _ := client.Do(request)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ErrorHTTP:", err)
	}

	resp.Body.Close()
	return string(data)
}
