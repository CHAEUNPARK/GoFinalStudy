package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
	Job     string `json:"job"`
}

func main() {
	var u User
	fo, err := os.Open("src/testcode/test.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fo.Close()

	dec := json.NewDecoder(fo)
	err = dec.Decode(&u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", u)

}
