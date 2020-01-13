package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Todo:전체적으로 예외처리 추가 필요
type MyConfig struct {
	// section
	Sections map[string]interface{}
	FileName string
}

func (app *MyConfig) Contains(str string, c string) bool {
	for _, value := range str {
		if c == string(value) {
			return true
		}
	}
	return false
}

func (app *MyConfig) Find(str string, c string) int {
	for i, value := range str {
		if c == string(value) {
			return i
		}
	}
	return len(str)
}

const whiteCharacter string = "\n\r \t"

func (app *MyConfig) removeWhiteSpace(line string) string {
	for {
		linechar := string(line[0])
		if app.Contains(whiteCharacter, linechar) {
			line = line[1:]
		} else {
			break
		}
	}

	for {
		linechar := string(line[len(line)-1])
		if app.Contains(whiteCharacter, linechar) {
			line = line[:len(line)-1]
		} else {
			break
		}
	}
	return line
}

func (app *MyConfig) Init(confFile string) (ret map[string]interface{}, err error) {
	// 1. 파일을 읽는다.
	fo, err := os.Open(confFile)
	if err != nil {
		return ret, err
	}
	defer fo.Close()

	// 2. config를 파싱한다.
	retM, err := app.Parse(fo)
	return retM, err
}

func (app *MyConfig) IsSection(line string) bool {
	var sBracketF string = "["
	var sBracketB string = "]"
	if check := app.Find(line, sBracketF); check == 0 {
		if check := app.Find(line, sBracketB); check == len(line)-1 {
			return true
		}
	}
	return false
}

func (app *MyConfig) parseSectionName(line string) (string, error) {
	sectionName := line[1 : len(line)-1]
	sectionName = app.removeWhiteSpace(sectionName)
	return sectionName, nil
}

func (app *MyConfig) Parse(fo *os.File) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	reader := bufio.NewReader(fo)
	var section map[string]interface{}
	var sectionName string
	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix {
			return ret, fmt.Errorf("byte로 담을 수 없는 길이입니다.")
		}
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println(err)
		}
		if len(line) == 0 {
			continue
		}

		// Remove White Space
		buff := app.removeWhiteSpace(string(line))

		// 섹션이 시작되었는가?
		//buff = string
		// Todo: value => string or int 구분 어떻게 할 것 인지
		if app.IsSection(buff) == true {
			// Parse Section name
			sectionName, err = app.parseSectionName(buff)
			if err != nil {
				//fmt.Println(err)
				return ret, err
			}
			// New Section
			section = make(map[string]interface{})
			ret[sectionName] = section
		} else if sectionName != "" {
			// Parse Identified
			key := strings.Split(buff, "=")[0]
			value := strings.Split(buff, "=")[1]

			key = app.removeWhiteSpace(key)
			value = app.removeWhiteSpace(value)
			host := ret[sectionName].(map[string]interface{})
			host[key] = value
		} else {
			continue
		}
	}
	return ret, nil
}

func (app *MyConfig) GetSection() {
}
func (app *MyConfig) GetParamInteger() {
}
func (app *MyConfig) GetParamString() {
}
func (app *MyConfig) GetParamBoolean() {
}
func main() {
	confFileName := "src/practice0109/config.conf"

	conf := MyConfig{}
	//err := conf.Init(confFileName)

	if _, err := conf.Init(confFileName); err != nil {
		fmt.Println(err)
		// error
		return
	} else {
		confRet, _ := conf.Init(confFileName)
		conf.FileName = confFileName
		conf.Sections = confRet
	}

	fmt.Println(conf)
	for key, value := range conf.Sections {
		fmt.Println(key, value)
	}
}
