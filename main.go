package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	// "github.com/serranoarevalo/learngo/dict"
)

var RequestFailed = errors.New("request failed")

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "wally"}
	for _, person := range people {
		go isSexy(person, c)
	}
	// Receiving a message is a blocking operation
	for i:=0; i < len(people); i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Waiting for messages")
}


func hitURL(url string) (error){
	fmt.Println("Checking:", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return RequestFailed
	}
	return nil
}

func isSexy(person string,  c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}