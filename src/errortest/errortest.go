package main

import (
	"fmt"
	"log"
	"time"
)

type HelloOneError struct {
	time  time.Time
	value int
}

func (e HelloOneError) Error() string {
	return fmt.Sprintf("%v: %d는 1이 아닙니다.", e.time, e.value)
}

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}
	return "", HelloOneError{time.Now(), n}
}
func main() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	s, err := helloOne(1)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(s)

	s, err = helloOne(2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(s)

	fmt.Println("Hello, world!")
}
