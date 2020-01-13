package main

import (
	"fmt"
	"os"
)

type SectionA struct {
	A int
	B int
}

type SectionB struct {
	C int
	D int
}

/*
func main() {

	var whitespace byte = ' '
	var sBracketF byte = '['
	var sBracketB byte = ']'
	var sectionName string
	secA := new(SectionA)
	secB := new(SectionB)
	reader := bufio.NewReader(fo)

	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix {
			fmt.Println("byte로 담을 수 없는 길이입니다.")
			return
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		if len(line) == 0 {
			continue
		}

		// Remove White Space
		buff := removeWhiteSpace(string(line))

		// 섹션이 시작되었는가?

		if IsSection(buff) == true {
			// Parse Section name
		} else {
			// Parse Identified
		}

		// >> IsSesion
		if check := Find(line, sBracketF); check == 0 {
			// removeWhiteSpace
			if check := Contains(line, whitespace); check {
				fmt.Println("Syntax Error : 공백이 들어가 있습니다.")
				return
			}

			// IsSession
			if check := Find(line, sBracketB); check != len(line)-1 {
				fmt.Println("Syntax Error : ] 가 없습니다.")
				return
			}
			sectionName = string(line[1 : len(line)-1])
		} else {
			if check := Contains(line, whitespace); check {
				fmt.Println("Syntax Error : 공백이 들어가 있습니다.")
				return
			}
			if sectionName == "SectionA" {
				key := strings.Split(string(line), "=")[0]
				if key == "A" {
					value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
					if err != nil {
						fmt.Println(err)
						return
					}
					secA.A = value
				} else if key == "B" {
					value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
					if err != nil {
						fmt.Println(err)
						return
					}
					secA.B = value
				} else {
					fmt.Println("Invalid Field Name : ", key)
					return
				}
			} else if sectionName == "SectionB" {
				key := strings.Split(string(line), "=")[0]
				if key == "C" { value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
					if err != nil {
						fmt.Println(err)
						return
					}
					secB.C = value
				} else if key == "D" {
					value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
					if err != nil {
						fmt.Println(err)
						return
					}
					secB.D = value
				} else {
					fmt.Println("Invalid Field Name : ", key)
					return
				}
			} else {
				fmt.Println("Invalid Section Name : ", sectionName)
				return
			}
		}
	}
	fmt.Println("SectionA")
	fmt.Println("A : ", secA.A)
	fmt.Println("B : ", secA.B)

	fmt.Println("SectionB")
	fmt.Println("C : ", secB.C)
	fmt.Println("D : ", secB.D)
}
*/

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
	// TRIM Left : 문자열의 왼쪽 화이트 스페이스를 제거한다.
	// TRIM Right : 문자열의 오른쪽 화이트 스페이스를 제거한다.
	// '    [    sectionA    ]   '
	// whitespace가 아닌 문자열 만날 때 까지 제거
	// 끝으로 가서 whitespace가 아닌 문자열 만날 때 까지 제거
	// 이건 parsing 함수에서
	// 대괄호 짝이 맞는지 체크
	// 대괄호 왼쪽 먼저 제거
	// 끝으로 가서 대괄호 오른쪽 제거
	// 대괄호 추출
	// section name이 아닐 경우 대괄호 체크 x
	// 값일 경우에도 똑같이 오른쪽 왼쪽만 제거
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

func (app *MyConfig) Init(confFile string) error {
	// TODO config set
	// 1. 파일을 읽는다.
	fo, err := os.Open(confFile)
	if err != nil {
		return err
	}
	defer fo.Close()

	// 2. config를 파싱한다.
	//return app.Parse(fo)
	return err
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

// TODO : return type
func (app *MyConfig) Parse(fo *os.File) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	//reader := bufio.NewReader(fo)
	//var section map[string]interface{}
	//var sectionName string
	//for {
	//	line, isPrefix, err := reader.ReadLine()
	//	if isPrefix {
	//		return ret, fmt.Errorf("byte로 담을 수 없는 길이입니다.")
	//	}
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//
	//		fmt.Println(err)
	//	}
	//	if len(line) == 0 {
	//		continue
	//	}
	//
	//	// Remove White Space
	//	buff := app.removeWhiteSpace(string(line))
	//
	//	// 섹션이 시작되었는가?
	//	//buff = string
	//	if IsSection(buff) == true {
	//		// Parse Section name
	//		sectionName = parseSectionName(buff)
	//		// New Section
	//		section = make(map[string]interface{})
	//		ret[sectionName] = section
	//	} else if sectionName != "" {
	//		// Parse Identified
	//	} else {
	//		continue
	//	}
	//
	//	// >> IsSesion
	//	if check := Find(line, sBracketF); check == 0 {
	//		// removeWhiteSpace
	//		if check := Contains(line, whitespace); check {
	//			fmt.Println("Syntax Error : 공백이 들어가 있습니다.")
	//			return
	//		}
	//
	//		// IsSession
	//		if check := Find(line, sBracketB); check != len(line)-1 {
	//			fmt.Println("Syntax Error : ] 가 없습니다.")
	//			return
	//		}
	//		sectionName = string(line[1 : len(line)-1])
	//	} else {
	//		if check := Contains(line, whitespace); check {
	//			fmt.Println("Syntax Error : 공백이 들어가 있습니다.")
	//			return
	//		}
	//		if sectionName == "SectionA" {
	//			key := strings.Split(string(line), "=")[0]
	//			if key == "A" {
	//				value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
	//				if err != nil {
	//					fmt.Println(err)
	//					return
	//				}
	//				secA.A = value
	//			} else if key == "B" {
	//				value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
	//				if err != nil {
	//					fmt.Println(err)
	//					return
	//				}
	//				secA.B = value
	//			} else {
	//				fmt.Println("Invalid Field Name : ", key)
	//				return
	//			}
	//		} else if sectionName == "SectionB" {
	//			key := strings.Split(string(line), "=")[0]
	//			if key == "C" {
	//				value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
	//				if err != nil {
	//					fmt.Println(err)
	//					return
	//				}
	//				secB.C = value
	//			} else if key == "D" {
	//				value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
	//				if err != nil {
	//					fmt.Println(err)
	//					return
	//				}
	//				secB.D = value
	//			} else {
	//				fmt.Println("Invalid Field Name : ", key)
	//				return
	//			}
	//		} else {
	//			fmt.Println("Invalid Section Name : ", sectionName)
	//			return
	//		}
	//	}
	//}
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
	//confFileName := "src/practice0109/config.conf"

	conf := MyConfig{}
	ret := conf.removeWhiteSpace("   2039 dfdf 838 	fsfsd\n\t\r")

	fmt.Println(ret)

	fmt.Println(conf.IsSection(ret))
	test := "[Sectionasdf]"
	fmt.Println(conf.Find(test,"["))
	fmt.Println(conf.Find(test, "]"))
	fmt.Println(conf.IsSection(test))
	//if err := conf.Init(confFileName); err != nil {
	//	// error
	//	return
	//} else {
	//}

	//
}
