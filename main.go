package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	// "github.com/serranoarevalo/learngo/dict"
)

var RequestFailed = errors.New("Request Failed")

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)	
}


func hitURL(url string) (error){
	fmt.Println("Checking:", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return RequestFailed
	}
	return nil
}

func isSexy(person string,  c chan bool) {
	fmt.Println(person)
	time.Sleep(time.Second * 5)
	c <- true
}