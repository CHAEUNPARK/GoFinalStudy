package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Section struct {
	name       string
	SectionVal []SectionVal
}

type SectionVal struct {
	name string
	val  int
}

func fileRead(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	sections := strings.Split(string(data), "[")
	sections = sections[1:]

	results := make([]Section, len(sections))
	for i, value := range sections {
		sec := new(Section)
		sectionName := strings.Split(value, "]")[0]
		sec.name = sectionName
		strValues := strings.Split(value, "]")[1]
		strValues = strings.TrimSpace(strValues)
		sectionValues := strings.Split(strValues, "\r\n")
		for _, value := range sectionValues {
			key := strings.Split(value, "=")[0]
			val := strings.Split(value, "=")[1]
			valInt, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
				return
			}
			var sVal = SectionVal{name: key, val: valInt}
			sec.SectionVal = append(sec.SectionVal, sVal)
		}
		results[i] = *sec
	}

	Print(results)

}

func Print(results []Section) {
	//fmt.Println(results)
	for _, value := range results {
		fmt.Println("section name : ", value.name)
		for _, value := range value.SectionVal {
			fmt.Println("key : ", value.name)
			fmt.Println("value : ", value.val)
		}
	}
}

func main() {
	fileRead("src/practice/config.conf")

}
