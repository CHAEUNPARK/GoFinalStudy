package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
		for _, cValue := range c{
			if cValue == value {
				return true
			}
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
	var c string
	for {
		c = string(line[0])
		if app.Contains(c, whiteCharacter) {
			line = line[1:]
		} else {
			break
		}
	}

	for {
		c = string(line[len(line)-1])
		if app.Contains(c, whiteCharacter) {
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
	ret, err = app.Parse(fo)
	return ret, err
}

func (app *MyConfig) IsSection(line string) (bool, error) {
	var sBracketF string = "["
	var sBracketB string = "]"
	if check := app.Find(line, sBracketF); check == 0 {
		if check := app.Find(line, sBracketB); check == len(line)-1 {
			return true, nil
		}
		return false, fmt.Errorf(line + " : ]가 없습니다.")
	}
	return false, nil
}

func (app *MyConfig) parseSectionName(line string) (string, error) {
	sectionName := line[1 : len(line)-1]
	sectionName = app.removeWhiteSpace(sectionName)
	if check := app.Contains(sectionName, whiteCharacter); check{
		return sectionName, fmt.Errorf(sectionName+" : Section name에 공백이 들어가 있습니다.")
	}
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
			return ret, fmt.Errorf(string(line) + " : byte로 담을 수 없는 길이입니다.")
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return ret, err
		}
		if len(line) == 0 {
			continue
		}

		// Remove White Space
		buff := app.removeWhiteSpace(string(line))

		// 섹션이 시작되었는가?
		//buff = string
		if _, err := app.IsSection(buff); err != nil {
			return ret, err
		} else if check, _ := app.IsSection(buff); check == true {
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

			host, ok := ret[sectionName].(map[string]interface{})
			if !ok {
				return ret, fmt.Errorf("asdf")
			}
			if _, err := strconv.Atoi(value); err != nil {
				host[key] = value
			} else {
				host[key], _ = strconv.Atoi(value)
			}
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
	confFileName := "src/practice0113/config.conf"

	conf := MyConfig{}
	confRet, err := conf.Init(confFileName)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		conf.FileName = confFileName
		conf.Sections = confRet
	}

	fmt.Println(conf)
	for key, value := range conf.Sections {
		fmt.Println(key, value)
	}
}
