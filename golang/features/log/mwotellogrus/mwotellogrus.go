package main

import (
	otellog "github.com/middleware-labs/golang-apm/mwotellogrus"
	track "github.com/middleware-labs/golang-apm/tracker"
	log "github.com/sirupsen/logrus"
)

func main() {

	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
	)

	logHook := otellog.NewMWOTelHook(config, otellog.WithLevels(log.AllLevels), otellog.WithName("otellogrus"))

	// add hook in logrus
	log.AddHook(logHook)
	// set formatter
	log.SetFormatter(&log.JSONFormatter{})
	// start logs
	log.Info("Info helloHandler testing 123456")
	log.Debug("Debug helloHandler testing 123456")
	log.Error("Error helloHandler testing 123456")

}
