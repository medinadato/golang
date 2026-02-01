package main

import (
	"fmt"
	"net/http"
)

func loadWebsites() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a new channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// function literal
		fmt.Println(l)
	}
	// time.Sleep(5 * time.Second)

	// for i := 0; i < len(links); i++ {
	// for l := range c {
	// 	// function literal
	// 	go func(link string) {
	// 		time.Sleep(2 * time.Second)
	// 		checkLink(link, c)
	// 	}(l)
	// }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- (link + " might be down!")
		return
	}

	c <- (link + " is up!")
}
