package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type SectionA struct {
	A, B int
}

func (sec SectionA) Print() {
	fmt.Println("SectionA")
	fmt.Println("A : ", sec.A, reflect.TypeOf(sec.A))
	fmt.Println("B : ", sec.B, reflect.TypeOf(sec.B))
}

type SectionB struct {
	C string
	D int
}

func (sec SectionB) Print() {
	fmt.Println("SectionB")
	fmt.Println("C : ", sec.C, reflect.TypeOf(sec.C))
	fmt.Println("D : ", sec.D, reflect.TypeOf(sec.D))
}

type SectionC struct {
	E, F string
}

func (sec SectionC) Print() {
	fmt.Println("SectionC")
	fmt.Println("E : ", sec.E, reflect.TypeOf(sec.E))
	fmt.Println("F : ", sec.F, reflect.TypeOf(sec.F))
}

func fileRead(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	sections := strings.Split(string(data), "[")
	sections = sections[1:]
	pArr := make([]Printer, len(sections))
	for i, value := range sections {
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
			pArr[i] = sectionA
		} else if sectionName == "SectionB" {
			sectionB := new(SectionB)
			sectionB.C = strings.Split(sectionValues[0], "=")[1]
			sectionB.D, err = strconv.Atoi(strings.Split(sectionValues[1], "=")[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			pArr[i] = sectionB
		} else if sectionName == "SectionC" {
			sectionC := new(SectionC)
			sectionC.E = strings.Split(sectionValues[0], "=")[1]
			sectionC.F = strings.Split(sectionValues[1], "=")[1]
			pArr[i] = sectionC
		}
	}
	for _, value := range pArr {
		value.Print()
	}

}

//
//func typeof(v interface{}) string {
//	return reflect.TypeOf(v).String()
//}
//
//func parsing(str string) []reflect.Value {
//	data := strings.Split(str, "=")[1]
//	typeOfData := typeof(data)
//
//	switch typeOfData {
//	case "int":
//		result, _ := strconv.Atoi(data)
//		return []reflect.Value{reflect.ValueOf(result)}
//	default:
//		return []reflect.Value{reflect.ValueOf(data)}
//	}
//}

type Printer interface {
	Print()
}

type Parser interface {
}

func main() {
	fileRead("src/configparsing/config.conf")
	//secA := new(SectionA)
	//a := parsing("A=\"1\"")
	//fmt.Println(a[0])
	//fmt.Println(reflect.TypeOf(a[0]))
	//
	//b := "asdf"
	//fmt.Println(reflect.TypeOf(b))

}
