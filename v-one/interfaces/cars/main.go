package main

import "fmt"

type car interface {
	getBrand() string
	getColor() string
}

type focus struct{}
type sportage struct{}
type comodore struct{}

func main() {
	f := focus{}
	printBrand(f)

	c := comodore{}
	printBrand(c)
}

func (f focus) getBrand() string {
	return "Ford"
}

func (s sportage) getBrand() string {
	return "Kia"
}

func (c comodore) getBrand() string {
	return "GM"
}

func (f focus) getColor() string {
	return "Blue"
}

func (s sportage) getColor() string {
	return "White"
}

func (c comodore) getColor() string {
	return "Red"
}

func printBrand(c car) {
	fmt.Println(c.getBrand())
}
