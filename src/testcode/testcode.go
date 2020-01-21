package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Formatter.(*logrus.TextFormatter).DisableColors = true
	log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout

}
