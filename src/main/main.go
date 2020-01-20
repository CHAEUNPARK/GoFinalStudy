package main

import (
	"MyLog"
)

func main() {
	logger := MyLog.MyLog{}
	logger.Init("src/logger/", "logger.log", "INFO")
	for i := 0 ; i < 10 ; i++{
		logger.WriteLogTrace(i, "hihi", i+1)
	}

}
