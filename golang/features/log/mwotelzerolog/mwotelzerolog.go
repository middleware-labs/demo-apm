package main

import (
	"github.com/middleware-labs/golang-apm/mwotelzerolog"
	track "github.com/middleware-labs/golang-apm/tracker"
	"github.com/rs/zerolog/log"
)

func main() {

	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
	)
	hook := mwotelzerolog.NewMWOTelHook(config)
	logger := log.Hook(hook)

	for {
		logger.Debug().Msg("Debug")
		logger.Info().Msg("Info")
		logger.Error().Msg("Error")
	}

}
