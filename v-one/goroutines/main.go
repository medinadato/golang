package goroutines

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(5 * time.Second)
}

func sayOla() {
	fmt.Println("Ola")
	// time.Sleep(*5time.Second)
}

func main() {
	go sayHello() // runs concurrently
	go sayOla()
	time.Sleep(10 * time.Second)
}
