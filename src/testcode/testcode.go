package main

import (
	"fmt"
	"reflect"
)

func main() {
	value := 1
	fmt.Println(reflect.TypeOf(value))
	str := fmt.Sprintln(value)
	fmt.Println(reflect.TypeOf(str))
	valueB := true
	fmt.Println(reflect.TypeOf(valueB))
	str = fmt.Sprintln(valueB)
	fmt.Println(reflect.TypeOf(str))

}
