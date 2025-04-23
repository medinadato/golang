package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func OnPage(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func main() {
	sum := 1
	for sum < 1000 {
		fmt.Println(OnPage("https://www.softwareexpress.com.br"), fmt.Sprintf("Call no. %d", sum))
		sum += 1
	}
}
