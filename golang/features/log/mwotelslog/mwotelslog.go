package main

import (
	"log"
	"log/slog"

	"github.com/middleware-labs/golang-apm/mwotelslog"
	track "github.com/middleware-labs/golang-apm/tracker"
)

func main() {

	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
	)

	logger := mwotelslog.NewMWOTelLogger(
		config,
		mwotelslog.WithDefaultConsoleLog(), // to enable console log
		mwotelslog.WithName("otelslog"),
	)

	//configure default logger
	slog.SetDefault(logger)

	logger.Debug("Debug")
	logger.Info("Info")
	logger.Error("Error")

	//Setting slog with default will also applies for standard log library.
	//With default log level
	log.Println("Debug standard")
	log.Println("Info standard")
	log.Println("Error standard")
}
