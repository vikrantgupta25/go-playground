package main

import "fmt"

func main() {

	// all the keys should be of same value and all the values should be of same value
	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"green": "#745",
	// }

	// var colors map[string]string

	colors := make(map[int]string)

	colors[10] = "#ffffff"
	fmt.Println(colors)

	// how to delete a key in the map
	delete(colors, 10)
	fmt.Println(colors)
}
