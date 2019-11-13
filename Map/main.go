package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":  "#ff0000",
		"blue": "#432192",
	}

	colors["white"] = "#ffffff"

	delete(colors, "red")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("color: ", color, ", hex: ", hex)
	}
}
