package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
func main() {
	fo, err := os.OpenFile(
		"src/logrustest/logger.log",
		os.O_CREATE|os.O_APPEND,
		644,
		)
	if err !=nil{
		fmt.Println(err)
		return
	}
	defer fo.Close()
	log.WithFields(log.Fields{
		"animal":"walrus",
		"size":10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":true,
		"number":122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":true,
		"number":100,
	}).Fatal("The ice breaks!")

}
