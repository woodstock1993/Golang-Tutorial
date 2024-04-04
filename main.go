package main

import (
	"errors"
	"fmt"
	"net/http"
	// "github.com/serranoarevalo/learngo/dict"
)

var RequestFailed = errors.New("Request Failed")

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.redit.com/",
		"https://www.soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",		
	}
	
	for _, value := range urls {
		result := "OK"
		err := hitURL(value)
		if err != nil {
			result = "FAILED"
		}
		results[value] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}	
}


func hitURL(url string) (error){
	fmt.Println("Checking:", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return RequestFailed
	}
	return nil
}