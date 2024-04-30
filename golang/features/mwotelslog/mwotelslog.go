package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/middleware-labs/golang-apm/mwotelslog"
	track "github.com/middleware-labs/golang-apm/tracker"
	sm "github.com/samber/slog-multi"
)

func main() {

	config, _ := track.Track(
		track.WithConfigTag("service", "your service name"),
		track.WithConfigTag("projectName", "your project name"),
	)

	logger := slog.New(
		//use slog-multi for if logging in console as well.
		sm.Fanout(
			slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}),
			mwotelslog.NewMWOtelHandler(config, mwotelslog.HandlerOptions{}),
		),
	)
	//configure default logger
	slog.SetDefault(logger)

	logger.Debug("Debug")
	logger.Info("Info")
	logger.Error("Error")

	//Setting slog with default will also applies for standard log library.
	//With default log level
	log.Println("Debug")
	log.Println("Info")
	log.Println("Error")
}
