package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetReportCaller(true)

	fo, err := os.OpenFile(
		"src/main/logrus.log",
		os.O_APPEND|os.O_CREATE,
		644,
	)
	if err != nil {
		log.Fatalln("file open fail")
	}
	log.SetOutput(fo)
	defer fo.Close()

	log.Infoln("infolevel")

}
