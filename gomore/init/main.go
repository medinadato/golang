package main

import "fmt"

func main() {
	fmt.Printf("Renato Go Go %s", "john \n")
	fmt.Println("3 + 5 = ", 3+5)
	fmt.Println(true && true)

	// variables
	a := "A value"
	fmt.Println("A = ", a)

	var b = "B value"
	fmt.Println("B = ", b)

	var c, d int = 1, 2
	fmt.Printf("C = %d, D = %d \n", c, d)

	var e int
	fmt.Println(e)

	var f string
	fmt.Println("\"", f, "\"")

	// constants
	const g string = "constant"
	const h = 100
	fmt.Println("G = ", g)
	fmt.Println(int(h))
	fmt.Println(int64(h))

	// for
	i := 0
	for i <= 5 {
		fmt.Println("i = ", i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println("j = ", j)
	}

	for i := range 3 {
		fmt.Println("i range ", i)
	}
}
