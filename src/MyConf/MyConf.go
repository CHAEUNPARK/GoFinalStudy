package Myconf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const whiteCharacter string = "\n\r \t"
const sBracketF string = "["
const sBracketB string = "]"
const newLine string = "\n"

// Todo:전체적으로 예외처리 추가 필요
type MyConfig struct {
	// section
	Sections map[string]map[string]string
	FileName string
}

func (app *MyConfig) contains(str string, c string) bool {
	for _, value := range str {
		for _, cValue := range c {
			if cValue == value {
				return true
			}
		}
	}
	return false
}

func (app *MyConfig) find(str string, c string) int {
	for i, value := range str {
		if c == string(value) {
			return i
		}
	}
	return len(str)
}

func (app *MyConfig) leftTrim(str string) string {
	for i, value := range str {
		isWhiteChar := false
		for _, cValue := range whiteCharacter {
			if cValue == value {
				isWhiteChar = true
				break
			}
		}

		if isWhiteChar == false {
			return str[i:]
		}
	}
	return str
}

func (app *MyConfig) rightTrim(str string) string {
	size := len(str)
	for i := size - 1; i >= 0; i-- {
		isWhiteChar := false
		for j := 0; j < len(whiteCharacter); j++ {
			if str[i] == whiteCharacter[j] {
				isWhiteChar = true
				break
			}
		}

		if isWhiteChar == false {
			return str[:i+1]
		}
	}
	return str
}

func (app *MyConfig) removeWhiteSpace(line string) string {
	s := app.leftTrim(line)
	return app.rightTrim(s)
}

func (app *MyConfig) Init(confFile string) error {
	// 1. 파일을 읽는다.
	fo, err := os.Open(confFile)
	if err != nil {
		return err
	}
	defer fo.Close()
	app.FileName = confFile
	// 2. config를 파싱한다.
	app.Sections, err = app.parse(fo)
	if err != nil {
		return err
	}

	return err
}

func (app *MyConfig) isSection(line string) (bool, error) {

	if line[0] == '[' && line[len(line)-1] == ']' {
		return true, nil
	} else if line[0] == '[' {
		return false, fmt.Errorf(line + " : ]가 없습니다.")
	} else if line[len(line)-1] == ']' {
		return false, fmt.Errorf(line + " : ]가 없습니다.")
	} else {
		return false, nil
	}

	return false, nil
}

func (app *MyConfig) parseSectionName(line string) (string, error) {
	sectionName := app.removeWhiteSpace(line[1 : len(line)-1])
	if check := app.contains(sectionName, whiteCharacter); check {
		return sectionName, fmt.Errorf(sectionName + " : Section name에 공백이 들어가 있습니다.")
	} else if check := app.contains(sectionName, sBracketF+sBracketB); check {
		return sectionName, fmt.Errorf(sectionName + " : Section name에 유효하지 않은 문자가 들어가 있습니다.")
	}
	return sectionName, nil
}

func (app *MyConfig) parse(fo *os.File) (map[string]map[string]string, error) {
	ret := make(map[string]map[string]string)
	reader := bufio.NewReader(fo)
	var section map[string]string
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

		// Remove White Space
		buff := app.removeWhiteSpace(string(line))

		if len(buff) == 0 {
			continue
		}
		// 섹션이 시작되었는가?
		//buff = string
		if check, err := app.isSection(buff); err != nil {
			return ret, err
		} else if check == true {
			// Parse Section name
			sectionName, err = app.parseSectionName(buff)
			if err != nil {
				return ret, err
			}
			// New Section
			section = make(map[string]string)
			ret[sectionName] = section
		} else if sectionName != "" {
			// Parse Identified
			data := strings.Split(buff, "=")
			key := data[0]
			value := data[1]

			key = app.removeWhiteSpace(key)
			value = app.removeWhiteSpace(value)

			value, err = app.valueCheck(value)
			if err != nil {
				return ret, err
			}
			ret[sectionName][key] = value

		} else {
			continue
		}
	}
	return ret, nil
}

func (app *MyConfig) sectionCheck(section string) (ret map[string]string, err error) {
	host, ok := app.Sections[section]
	if !ok {
		return ret, fmt.Errorf("There is no section name : " + section)
	}
	return host, nil
}

func (app *MyConfig) valueCheck(value string) (ret string, err error) {
	if _, err := strconv.Atoi(value); err != nil {
		if _, err = strconv.ParseBool(value); err != nil {
			if value[0] == '"' && value[len(value)-1] == '"' {
				return value[1 : len(value)-1], nil
			} else {
				return ret, fmt.Errorf("Invalid Syntax : " + value)
			}
		}
	}
	return value, nil
}

func (app *MyConfig) parseValue(section string, param string) (ret string, err error) {
	host, err := app.sectionCheck(section)
	if err != nil {
		return ret, err
	}

	value, ok := host[param]
	if !ok {
		return ret, fmt.Errorf("There is no key name : " + param)
	}
	return value, nil
}

//Todo: file write
//Section 존재할 경우 그밑에 쓰기
//Section 존재하지 않는 경우 맨 밑에 쓰기
//파일이 변경되었을때
//modified time
func (app *MyConfig) writeConfig() {
	fo, err := os.Create(app.FileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer fo.Close()
	str := ""
	sectionNames, err := app.GetSectionList()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for _, value := range sectionNames {
		str = str + "[" + value + "]" + newLine
		params, err := app.GetSection(value)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		for key, value := range params {
			if _, err := strconv.Atoi(value); err != nil {
				if _, err := strconv.ParseBool(value); err != nil {
					value = "\"" + value + "\""
				}
			}
			str = str + key + " = " + value + newLine
		}
	}
	_, err = fo.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("set param completely")
}

//Todo: 사용자
func (app *MyConfig) GetSectionList() (ret []string, err error) {
	if len(app.Sections) == 0 {
		return ret, fmt.Errorf("No sections")
	}
	for key, _ := range app.Sections {
		ret = append(ret, key)
	}
	return ret, nil
}

func (app *MyConfig) GetSection(section string) (ret map[string]string, err error) {
	host, err := app.sectionCheck(section)
	if err != nil {
		return ret, err
	}
	ret = map[string]string{}
	if len(host) != 0 {
		return host, nil
	} else {
		return ret, fmt.Errorf("Empty Section")
	}

}

func (app *MyConfig) GetParamInteger(section string, param string) (ret int, err error) {
	value, err := app.parseValue(section, param)
	if err != nil {
		return ret, err
	}
	// map[string]string

	if ret, err = strconv.Atoi(value); err != nil {
		return ret, err
	} else {
		return ret, nil
	}
}

func (app *MyConfig) GetParamString(section string, param string) (ret string, err error) {
	value, err := app.parseValue(section, param)
	if err != nil {
		return ret, err
	}

	return value, nil
}

func (app *MyConfig) GetParamBoolean(section string, param string) (ret bool, err error) {
	value, err := app.parseValue(section, param)
	if err != nil {
		return ret, err
	}

	if ret, err = strconv.ParseBool(value); err != nil {
		return ret, err
	} else {
		return ret, nil
	}
}

func (app *MyConfig) SetParamInteger(section string, key string, value int) {
	if app.contains(section, whiteCharacter+sBracketB+sBracketF) {
		fmt.Println("Invalid Section Name : " + section)
		os.Exit(0)
	}
	host, ok := app.Sections[section]
	if !ok {
		newSection := make(map[string]string)
		newSection[key] = strconv.Itoa(value)
		app.Sections[section] = newSection
	} else {
		host[key] = strconv.Itoa(value)
	}
	app.writeConfig()
}

func (app *MyConfig) SetParamString(section string, key string, value string) {
	if app.contains(section, whiteCharacter+sBracketB+sBracketF) {
		fmt.Println("Invalid Section Name : " + section)
		os.Exit(0)
	}
	host, ok := app.Sections[section]
	if !ok {
		newSection := make(map[string]string)
		newSection[key] = value
		app.Sections[section] = newSection
	} else {
		host[key] = value
	}
	app.writeConfig()
}

func (app *MyConfig) SetParamBoolean(section string, key string, value bool) {
	if app.contains(section, whiteCharacter+sBracketB+sBracketF) {
		fmt.Println("Invalid Section Name : " + section)
		os.Exit(0)
	}
	host, ok := app.Sections[section]
	if !ok {
		newSection := make(map[string]string)
		newSection[key] = strconv.FormatBool(value)
		app.Sections[section] = newSection
	} else {
		host[key] = strconv.FormatBool(value)
	}
	app.writeConfig()
}
