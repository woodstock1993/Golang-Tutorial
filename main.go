package main

import (
	"fmt"
	"github.com/serranoarevalo/learngo/dict"
)

func main() {
	dictionary := dict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	
	// Add 하지 못했을 때
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("key", word, "value:", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	// }
	}
}