package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
	Job     string `json:"job"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	rate := math.Round((rand.Float64()*100)/0.005) * 0.005
	fmt.Println(rate)
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64() * 100)
	fmt.Println(rand.Float64() * 100 / 0.005)
	fmt.Println((rand.Float64() * 100 / 0.005) * 0.005)
}
