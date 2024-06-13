package main

import (
	"os"

	mwotelzap "github.com/middleware-labs/golang-apm/mwotelzap"
	track "github.com/middleware-labs/golang-apm/tracker"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
	)

	// Define encoder configuration
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Configure console output
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)
	consoleLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleDebugging, consoleLevel)

	// Configure file output
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}
	defer file.Close()
	fileEncoder := zapcore.NewConsoleEncoder(encoderCfg)  // You can use a different encoder if needed
	fileLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel) // Set your desired log level here
	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(file), fileLevel)

	logger := zap.New(zapcore.NewTee(consoleCore, fileCore, mwotelzap.NewMWOTelCore(config, mwotelzap.WithName("otelzaplog"))))
	zap.ReplaceGlobals(logger)

	//logs
	logger.Error("Error zap")
	logger.Info("Info zap")
	logger.Warn("Warn zap")

}
