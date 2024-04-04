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
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 5)
}


func hitURL(url string) (error){
	fmt.Println("Checking:", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return RequestFailed
	}
	return nil
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}