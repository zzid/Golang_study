package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)
type requestResult struct {
	url 	string
	status 	string
}
var errRequestFailed = errors.New("Request Failed")


// specify what channel can do in here
func hitURL(url string, c chan<- requestResult) {
	// fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400{
		c <- requestResult{url: url, status: "FAILED"}
		// fmt.Println(err, resp.StatusCode)
		// return errRequestFailed
	}
	c <- requestResult{url: url, status: status}
}

func main() {
	/* 
		url checker
	*/
	results := make(map[string] string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://github.com/zzid",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i<len(urls); i++ {
		result := <- c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
	/*
		Go routines
			goroutines is alive while main is alive

		Channels

	*/
	// // This will work
	// // go sexyCount("zzid")
	// // sexyCount("Nadia")
	
	// // This will nothing
	// // go sexyCount("zzid")
	// // go sexyCount("Nadia")
	// c := make(chan string)
	// people := []string{"zzid", "nadia", "they", "are", "in love"}
	// for _, person := range people {
	// 	go isSexy(person, c)
	// }
	// // result := <-c
	// // fmt.Println(result)
	// // fmt.Println(<-c)
	// // fmt.Println(<-c)

	// /* receiving is "blocking operation" */
	// for i :=0; i<len(people); i++ {
	// 	fmt.Println(<-c)
	// }
}



func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	channel <- person + " is sexy"
}


func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}