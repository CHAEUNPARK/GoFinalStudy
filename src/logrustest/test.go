package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)
//logrus같은 경우에는 rotation 기능이 없기 때문에 rotation 기능은 hook을 이용해서 씀(보통)
func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
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
	log.SetOutput(fo)
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
