package main

import (
	"math/rand/v2"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	g "github.com/middleware-labs/golang-apm-gin/gin"
	mwotelzap "github.com/middleware-labs/golang-apm/mwotelzap"
	track "github.com/middleware-labs/golang-apm/tracker"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	config, _ := track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
		track.WithConfigTag(track.Token, "your token"),
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

	logger := zap.New(zapcore.NewTee(consoleCore, fileCore, mwotelzap.NewMWOTelCore(config)))
	zap.ReplaceGlobals(logger)

	go makeRequest(logger)

	r := gin.Default()
	r.Use(g.Middleware(config))
	r.GET("/books", FindBooks)
	r.Run(":8090")
}

func FindBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func makeRequest(logger *zap.Logger) {
	for {
		logger.Error("Error")
		logger.Info("Info")
		logger.Warn("Warn")
		time.Sleep(time.Duration(rand.IntN(100)) * 10 * time.Millisecond)
		http.Get("http://localhost:8090/books")
	}
}
