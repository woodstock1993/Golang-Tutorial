package main

import (
	"fmt"
	"github.com/serranoarevalo/learngo/dict"
)

func main() {
	dic := dict.Dictionary{"first": "First word"}
	definition, err := dic.Search("first")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(definition)
}
