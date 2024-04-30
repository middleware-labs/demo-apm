package main

import (
	"github.com/middleware-labs/golang-apm/mwotelzerolog"
	track "github.com/middleware-labs/golang-apm/tracker"
	"github.com/rs/zerolog/log"
)

func main() {

	config, _ := track.Track(
		track.WithConfigTag("service", "your service name"),
		track.WithConfigTag("projectName", "your project name"),
	)
	hook := mwotelzerolog.NewMWOtelHook(config)
	logger := log.Hook(hook)

	logger.Debug().Msg("Debug")
	logger.Info().Msg("Info")
	logger.Error().Msg("Error")

}
