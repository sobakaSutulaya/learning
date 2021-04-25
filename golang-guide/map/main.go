package main

import "fmt"

func main() {
	//few way to init map
	//first
	colors := map[string]string{
		"red": "#ff0000",
		"white" : "#ffffff",
		"black" : "#000000",
	}

	//second through var
	var varColors map[string]string

	//third through make
	makeColors := make(map[string] string)
	makeColors["white"] = "#fffffff"

	//delete by key
	delete(makeColors, "white")

	printMap(colors)

	fmt.Println(colors)
	fmt.Println(varColors)
	fmt.Println(makeColors)
}

func printMap(c map[string]string) {
	//iterate through map
	for color, hex := range c {
		fmt.Println("Hec code for", color, "is", hex)
	}
}