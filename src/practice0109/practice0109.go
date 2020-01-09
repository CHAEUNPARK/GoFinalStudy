package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

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

		if check := Find(line, sBracketF); check == 0 {
			if check := Contains(line, whitespace); check {
				fmt.Println("Syntax Error : 공백이 들어가 있습니다.")
				break
			}
			if check := Find(line, sBracketB); check != len(line)-1 {
				fmt.Println("Syntax Error : ] 가 없습니다.")
				break
			}
			sectionName := string(line[1 : len(line)-1])
			fmt.Println(sectionName)
		} else {
			if check := Contains(line, whitespace); check {
				fmt.Println("Syntax Error : 공백이 들어가 있습니다.")
				break
			}
			key := strings.Split(string(line), "=")[0]
			value := strings.Split(string(line), "=")[1]
			fmt.Println(key, value)
		}

	}
}
