package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name  string `json:"name"`
	Hobby string
	age   int
}

func main() {
	p := person{
		Name:  "Amy",
		age:   18,
		Hobby: "唱歌",
	}
	bytes, _ := json.Marshal(&p)
	fmt.Println(string(bytes))
}
