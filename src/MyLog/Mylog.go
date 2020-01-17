package MyLog

import (
	Myconf "MyConf"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

const (
	trace = iota
	debug
	info
	warn
	errC
	fatal
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
		app.LEVEL = trace
	case "DEBUG":
		app.LEVEL = debug
	case "INFO":
		app.LEVEL = info
	case "WARN":
		app.LEVEL = warn
	case "ERROR":
		app.LEVEL = errC
	case "FATAL":
		app.LEVEL = fatal
	}
	return nil
}
func (app *MyLog) write(log string, level string) {
	//Todo : file open -> 덮어쓰지 않고 밑으로 추가
	if _, err := os.Stat(app.PATH); os.IsNotExist(err){
		err = os.Mkdir(app.PATH, 644)
		if err != nil{
			fmt.Println(err)
			return
		}
	}
	fo, err := os.OpenFile(
		app.PATH+app.FILENAME,
		os.O_CREATE|os.O_APPEND,
		644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fo.Close()
	//date time
	datetime := time.Now().Format("2006-01-02 15:04:05")
	_, fileName, lineNo := app.getInfo()
	_, fileName = path.Split(fileName)

	str := "[" + datetime + "] [" + level + "] [" + fileName + "] [line : " + strconv.Itoa(lineNo) + "] " + log + "\n"
	fo.Write([]byte(str))
	return
}
func (app *MyLog) WriteLogTrace(log string) {
	if app.LEVEL < trace {
		return
	}
	app.write(log, "Trace")
}

func (app *MyLog) WriteLogDebug(log string) {
	if app.LEVEL < debug {
		return
	}
	app.write(log, "Debug")
}

func (app *MyLog) WriteLogInfo(log string) {
	if app.LEVEL < info {
		return
	}
	app.write(log, "Information")
}

func (app *MyLog) WriteLogWarn(log string) {
	if app.LEVEL < warn {
		return
	}
	app.write(log, "Warning")
}

func (app *MyLog) WriteLogError(log string) {
	if app.LEVEL < errC {
		return
	}
	app.write(log, "Error")
}

func (app *MyLog) WriteLogFatal(log string) {
	if app.LEVEL < fatal {
		return
	}
	app.write(log, "Fatal")
}

func (app *MyLog) getInfo() (string, string, int) {
	pc, fileName, lineNo, _ := runtime.Caller(3)
	funcName := runtime.FuncForPC(pc).Name()
	return funcName, fileName, lineNo
}

