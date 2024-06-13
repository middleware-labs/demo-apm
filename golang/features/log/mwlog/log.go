package main

import (
	"fmt"

	"github.com/middleware-labs/golang-apm/logger"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {
	track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
	)
	fmt.Println("121212")
	logger.Error("Error")
	logger.Info("Info")
	logger.Warn("Warn")
}
