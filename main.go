package main

import (
	"fmt"
	"github.com/serranoarevalo/learngo/dict"
)

func main() {
	dic := dict.Dictionary{}
	err := dic.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	value, _ := dic.Search("hello")
	fmt.Println(value)
	err2 := dic.Add("hello", "Greeting")
	if err2 != nil{
		fmt.Println(err2)
	}
}
