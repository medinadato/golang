package main

import "fmt"

func main() {
	// #1 way to start map
	// colors := map[string]string{
	// 	"red":  "#FF0000",
	// 	"green": "#4bf745",
	// }

	// #2 way to start a map
	// var colors map[string]string
	// colors = map[string]string{
	// 	"white": "#ffffff",
	// 	"black": "#000000",
	// }

	// #3 way to start a map
	colors := make(map[string]string)
	colors["red"] = "#FF0000"
	colors["white"] = "#ffffff"

	// testing deleting values
	// colors := make(map[int]string)
	// colors[7] = "#FF0000"
	// colors[99] = "#ffffff"
	// delete(colors, 7)
	// colors[64] = "#8h347f"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
