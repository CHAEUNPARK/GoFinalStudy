package main

import (
	"MyLog"
	"strconv"
)

func main() {
	logger := MyLog.MyLog{}
	logger.Init()
	for i := 0 ; i < 10 ; i++{
		logger.WriteLogDebug(strconv.Itoa(i))
	}

}
