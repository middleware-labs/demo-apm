package main

import (
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
		track.WithConfigTag(track.Token, "your access token"),
	)
}
