package main

import (
	Myconf "MyConf"
	"fmt"
	"time"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type MyLog struct {
	PATH     string
	FILENAME string
	LEVEL    int
}

func (app *MyLog) Init() error {
	conf := Myconf.MyConfig{}
	conf.Init("src/MyConf/config.conf")
	path, err := conf.GetParamString("LOG", "PATH")
	if err != nil {
		return err
	}
	filename, err := conf.GetParamString("LOG", "FILENAME")
	if err != nil {

		return err
	}
	level, err := conf.GetParamString("LOG", "LEVEL")
	if err != nil {
		return err
	}
	app.PATH = path
	app.FILENAME = filename
	switch level {
	case "TRACE":
		app.LEVEL = TRACE
	case "DEBUG":
		app.LEVEL = DEBUG
	case "INFO":
		app.LEVEL = INFO
	case "WARN":
		app.LEVEL = WARN
	case "ERROR":
		app.LEVEL = ERROR
	case "FATAL":
		app.LEVEL = FATAL
	}
	return nil
}
func (app *MyLog) Write(log string) {
	//Todo : file open -> 덮어쓰지 않고 밑으로 추가
	//date time
	//log level    ->  고려 사항 : 상수화 시켜서 크기 비교 하는게 나을지 string 그대로 if문 거는게 나을지
	//file path
	//line
	//msg
	return
}
func (app *MyLog) WriteLogTrace(log string) {
	if app.LEVEL > TRACE {
		return
	}
	app.Write(log)
}

func (app *MyLog) WriteLogDebug(log string) {
	if app.LEVEL > DEBUG {
		return
	}
	app.Write(log)
}

func (app *MyLog) WriteLogInfo(log string) {
	if app.LEVEL > INFO {
		return
	}
	app.Write(log)
}

func (app *MyLog) WriteLogWarn(log string) {
	if app.LEVEL > WARN{
		return
	}
	app.Write(log)
}

func (app *MyLog) WriteLogError(log string) {
	if app.LEVEL > ERROR {
		return
	}
	app.Write(log)
}

func (app *MyLog) WriteLogFatal(log string) {
	if app.LEVEL > FATAL {
		return
	}
	app.Write(log)
}

func main() {
	fmt.Println(TRACE, DEBUG, INFO, WARN, ERROR, FATAL)
	logger := MyLog{}
	logger.Init()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(logger.LEVEL)

}
