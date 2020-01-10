package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SectionA struct {
	A int
	B int
}

type SectionB struct {
	C int
	D int
}

func Contains(bs []byte, b byte) bool {
	for _, value := range bs {
		if b == value {
			return true
		}
	}
	return false
}

func Find(bs []byte, b byte) int {
	for i, value := range bs {
		if b == value {
			return i
		}
	}
	return len(bs)
}

func main() {
	fo, err := os.Open("src/practice0109/config.conf")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fo.Close()
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
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		if len(line) == 0{
			continue
		}
		if check := Find(line, sBracketF); check == 0 {
			if check := Contains(line, whitespace); check {
				fmt.Println("Syntax Error : 공백이 들어가 있습니다.")
				return
			}
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
				if key := strings.Split(string(line), "=")[0]; key == "A" {
					value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
					if err != nil {
						fmt.Println(err)
						return
					}
					secA.A = value
				} else if key := strings.Split(string(line), "=")[0]; key == "B" {
					value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
					if err != nil {
						fmt.Println(err)
						return
					}
					secA.B = value
				} else {
					fmt.Println("Invalid Section Name")
					return
				}
			} else if sectionName == "SectionB" {
				if key := strings.Split(string(line), "=")[0]; key == "C" {
					value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
					if err != nil {
						fmt.Println(err)
						return
					}
					secB.C = value
				} else if key := strings.Split(string(line), "=")[0]; key == "D" {
					value, err := strconv.Atoi(strings.Split(string(line), "=")[1])
					if err != nil {
						fmt.Println(err)
						return
					}
					secB.D = value
				} else {
					fmt.Println("Invalid Section Name")
					return
				}
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
