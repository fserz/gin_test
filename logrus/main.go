package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.Formatter = &logrus.JSONFormatter{}
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	f, _ := os.Create("./gin.log")
	log.Out = f
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	// Only log the warning severity or above.
	log.Level = logrus.InfoLevel
}

func main() {
	//log.SetReportCaller(true)

	log.WithFields(logrus.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")

	log.WithFields(logrus.Fields{
		"event": "event",
		"topic": "topic",
		"key":   "key",
	}).Fatal("Failed to send event")
}
