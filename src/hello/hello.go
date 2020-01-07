package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := []string{"Maria", "John", "Andrew", "James"}

	r := ring.New(len(data))
	for i := 0; i < r.Len(); i++ {
		r.Value = data[i]
		r = r.Next()
	}

	r.Do(func(x interface{}) {
		fmt.Println(x)
	})

	fmt.Println("Move forward :")
	r = r.Move(1)

	fmt.Println("curr : ", r.Value)
	fmt.Println("Prev : ", r.Prev().Value)
	fmt.Println("Next : ", r.Next().Value)

}
