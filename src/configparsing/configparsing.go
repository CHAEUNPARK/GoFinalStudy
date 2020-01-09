package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type SectionA struct {
	A int
	B int
}

type SectionB struct {
	C string
	D int
}

type SectionC struct {
	E string
	F string
}

func parsing(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	sections := strings.Split(string(data), "[")
	sections = sections[1:]
	for _, value := range sections {
		sectionName := strings.Split(value, "]")[0]
		strValues := strings.Split(value, "]")[1]
		strValues = strings.TrimSpace(strValues)
		sectionValues := strings.Split(strValues, "\r\n")
		if sectionName == "SectionA" {
			sectionA := new(SectionA)
			sectionA.A, err = strconv.Atoi(strings.Split(sectionValues[0], "=")[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			sectionA.B, err = strconv.Atoi(strings.Split(sectionValues[1], "=")[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(sectionA.A, sectionA.B)
		} else if sectionName == "SectionB" {
			sectionB := new(SectionB)
			sectionB.C = strings.Split(sectionValues[0], "=")[1]
			sectionB.D, err = strconv.Atoi(strings.Split(sectionValues[1], "=")[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(sectionB.C, sectionB.D)
		} else if sectionName == "SectionC" {
			sectionC := new(SectionC)
			sectionC.E = strings.Split(sectionValues[0], "=")[1]
			sectionC.F = strings.Split(sectionValues[1], "=")[1]
			fmt.Println(sectionC.E, sectionC.F)
		}
	}
}

func Print() {

}

func main() {
	parsing("src/configparsing/config.conf")
}
