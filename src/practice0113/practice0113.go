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
		for _, cValue := range c {
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

func (app *MyConfig) typeCheck(str string) (ret string, err error) {
	if _, err := strconv.Atoi(str); err != nil {
		if _, err := strconv.ParseBool(str); err != nil {
			if str[0] == '"' && str[len(str)-1] == '"' {
				return "string", nil
			} else if str[0] == '"' || str[len(str)-1] == '"' {
				return ret, fmt.Errorf("Invalid Syntax : " + str)
			} else {
				return "string", nil
			}
		} else {
			return "bool", nil
		}
	} else {
		return "int", nil
	}
}

const whiteCharacter string = "\n\r \t"
const sBracketF string = "["
const sBracketB string = "]"

func (app *MyConfig) LeftTrim(str string) string {
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

func (app *MyConfig) RightTrim(str string) string {
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
	s := app.LeftTrim(line)
	return app.RightTrim(s)
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
	app.Sections, err = app.Parse(fo)
	if err != nil {
		return err
	}

	return err
}

func (app *MyConfig) IsSection(line string) (bool, error) {

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
	if check := app.Contains(sectionName, whiteCharacter); check {
		return sectionName, fmt.Errorf(sectionName + " : Section name에 공백이 들어가 있습니다.")
	} else if check := app.Contains(sectionName, sBracketF+sBracketB); check {
		return sectionName, fmt.Errorf(sectionName + " : Section name에 유효하지 않은 문자가 들어가 있습니다.")
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

		// Remove White Space
		buff := app.removeWhiteSpace(string(line))

		if len(buff) == 0 {
			continue
		}
		// 섹션이 시작되었는가?
		//buff = string
		if check, err := app.IsSection(buff); err != nil {
			return ret, err
		} else if check == true {
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
			data := strings.Split(buff, "=")
			key := data[0]
			value := data[1]

			key = app.removeWhiteSpace(key)
			value = app.removeWhiteSpace(value)

			host := ret[sectionName].(map[string]interface{})
			//문자열 체크 ""
			//type check
			/*
				//int -> 숫자로만 이루어진 경우
				string -> ""가 있는 경우
				//bool -> true / false
			*/

			if _, err := strconv.Atoi(value); err != nil {
				if valBool, err := strconv.ParseBool(value); err != nil {
					if value[0] == '"' && value[len(value)-1] == '"' {
						host[key] = value[1 : len(value)-1]
					} else if value[0] == '"' || value[len(value)-1] == '"' {
						return ret, fmt.Errorf("Invalid Syntax : \"가 없습니다.(" + buff + ")")
					} else {
						host[key] = value
					}
				} else {
					host[key] = valBool
				}
			} else {
				host[key], _ = strconv.Atoi(value)
			}
		} else {
			continue
		}
	}
	return ret, nil
}

//Todo:작성
func (app *MyConfig) GetSectionList() (ret []string, err error) {
	if len(app.Sections) == 0 {
		return ret, fmt.Errorf("No sections")
	}
	for key, _ := range app.Sections {
		ret = append(ret, key)
	}
	return ret, nil
}

func (app *MyConfig) GetSection(section string) (ret map[string]interface{}, err error) {
	host, ok := app.Sections[section].(map[string]interface{})
	if !ok {
		//host = make(map[string]interface{})
		return ret, fmt.Errorf("There is no Section name : " + section)
	}
	ret = map[string]interface{}{}
	if len(host) != 0 {
		for key, value := range host {
			ret[key] = value
		}
	} else {
		return ret, fmt.Errorf("Empty Section")
	}

	return ret, nil
}

/*
func (app *MyConfig) GetValue(conf MyConfig, secName string, paramName string) (ret interface{}, err error){
	host := conf.Sections[secName].(map[string]interface{})
	switch host[paramName].(type) {
	case int:

	}
	return ret, nil
}
*/
func (app *MyConfig) IsCheck(section string, param string) (ret map[string]interface{}, err error) {
	host, ok := app.Sections[section].(map[string]interface{})
	if !ok{
		return ret, fmt.Errorf("There is no section name : " + section)
	}

	_, ok = host[param]
	if !ok{
		return ret, fmt.Errorf("There is no key name : "+param)
	}
	return host, nil
}

func (app *MyConfig) GetParamInteger(section string, param string) (ret int, err error) {
	//host, ok := app.Sections[section].(map[string]interface{})
	//if !ok {
	//	return ret, fmt.Errorf("There is no section name : " + section)
	//}
	//
	//_, ok = host[param]
	//if !ok {
	//	return ret, fmt.Errorf("There is no key name : " + param)
	//}
	host, err := app.IsCheck(section, param)

	value, ok := host[param].(int)
	if !ok {
		return ret, fmt.Errorf(section + "'s " + param + " is not int")
	}

	ret = value

	return ret, nil
}

func (app *MyConfig) GetParamString(section string, param string) (ret string, err error) {
	host, ok := app.Sections[section].(map[string]interface{})
	if !ok {
		return ret, fmt.Errorf("There is no section name : " + section)
	}

	_, ok = host[param]
	if !ok {
		return ret, fmt.Errorf("There is no key name : " + param)
	}

	value, ok := host[param].(string)
	if !ok {
		return ret, fmt.Errorf(section + "'s " + param + " is not string")
	}

	ret = value

	return ret, nil
}

func (app *MyConfig) GetParamBoolean(section string, param string) (ret bool, err error) {
	host, ok := app.Sections[section].(map[string]interface{})
	if !ok {
		return ret, fmt.Errorf("There is no section name : " + section)
	}

	_, ok = host[param]
	if !ok {
		return ret, fmt.Errorf("There is no key name : " + param)
	}

	value, ok := host[param].(bool)
	if !ok {
		return ret, fmt.Errorf(section + "'s " + param + " is not boolean")
	}

	ret = value

	return ret, nil
}

func (app *MyConfig) SetParamInteger(section string, param string) (ret int, err error) {
	return ret, nil
}

func (app *MyConfig) SetParamString(section string, param string) (ret string, err error) {
	return ret, nil
}

func (app *MyConfig) SetParamBoolean(section string, param string) (ret string, err error) {
	return ret, err
}

//func (app *MyConfig) WriteConfig() error
func main() {
	confFileName := "src/practice0113/config.conf"

	conf := MyConfig{}
	err := conf.Init(confFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(conf)
	sections, err := conf.GetSectionList()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sections)

	getInt, err := conf.GetParamInteger("SectionA", "A")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getInt)

	getStr, err := conf.GetParamString("SectionC", "E")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getStr)

	getBool, err := conf.GetParamBoolean("SectionD", "G")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getBool)

	sec, err := conf.GetSection("SectionD")
	fmt.Println(sec)

	confFileName = "src/practice0109/config.conf"

	conf = MyConfig{}
	err = conf.Init(confFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(conf)
	sections, err = conf.GetSectionList()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sections)

	getInt, err = conf.GetParamInteger("SectionA", "A")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getInt)

	sec, err = conf.GetSection("SectionD")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(sec)

	getStr, err = conf.GetParamString("SectionC", "E")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getStr)

	getBool, err = conf.GetParamBoolean("SectionD", "G")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getBool)



	//tc, err := conf.typeCheck("true")
	//fmt.Println(tc, err)
	//tc1, err := conf.typeCheck("1")
	//fmt.Println(tc1, err)
	//tc2, err := conf.typeCheck("\"dfssadf\"")
	//fmt.Println(tc2, err)
	//
	//tc3, err := conf.typeCheck("\"fdfdfdfdfd")
	//fmt.Println(tc3, err)

}
