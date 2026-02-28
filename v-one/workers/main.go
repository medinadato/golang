package main

import "fmt"

func square(n int, out chan<- int) {
	out <- n * n
}

func main() {
	out := make(chan int)

	for i := 1; i <= 3; i++ {
		go square(i, out)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println(<-out)
	}
}
