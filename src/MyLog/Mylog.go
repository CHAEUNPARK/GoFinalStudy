package MyLog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

//Todo: set level, parameter interface, getter&setter

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

func (app *MyLog) Init(path string, filename string, level string) error {
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
func (app *MyLog) write(msg string, level string) {
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

	str := "[" + datetime + "] [" + level + "] [" + fileName + "] [line : " + strconv.Itoa(lineNo) + "] " + msg + "\n"
	fo.Write([]byte(str))
	return
}
func (app *MyLog) WriteLogTrace(args ...interface{}) {
	if app.LEVEL < trace {
		return
	}
	msg := app.formatString(args)
	app.write(msg, "Trace")
}

func (app *MyLog) WriteLogDebug(args ...interface{}) {
	if app.LEVEL < debug {
		return
	}
	msg := app.formatString(args)
	app.write(msg, "Debug")
}

func (app *MyLog) WriteLogInfo(args ...interface{}) {
	if app.LEVEL < info {
		return
	}
	msg := app.formatString(args)
	app.write(msg, "Information")
}

func (app *MyLog) WriteLogWarn(args ...interface{}) {
	if app.LEVEL < warn {
		return
	}
	msg := app.formatString(args)
	app.write(msg, "Warning")
}

func (app *MyLog) WriteLogError(args ...interface{}) {
	if app.LEVEL < errC {
		return
	}
	msg := app.formatString(args)
	app.write(msg, "Error")
}

func (app *MyLog) WriteLogFatal(args ...interface{}) {
	if app.LEVEL < fatal {
		return
	}
	msg := app.formatString(args)
	app.write(msg, "Fatal")
}

func (app *MyLog) getInfo() (string, string, int) {
	pc, fileName, lineNo, _ := runtime.Caller(3)
	funcName := runtime.FuncForPC(pc).Name()
	return funcName, fileName, lineNo
}
func (app *MyLog) formatString(args []interface{}) string {
	var msg string
	for _, value := range args{
		msg += fmt.Sprint(value)
	}
	return msg
}

