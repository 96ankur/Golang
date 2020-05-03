package main

import (
	"fmt"
)

/*
* Map is a data structure just like map in solidity, but in map all the keys must be of SAME TYPE
* and silimarly all the values must be of SAME TYPE
*
* But we can iterate over map in golang	
*
*/
func main(){
	//  SYNTAX 1

 	colors := map[string]string{
		"red":"#ff0000",
		"green":"#523698",
		"white":"#258745",
	}

	//  SYNTAX 2

	// var colors map[string]string

	//  SYNTAX 3

	// colors := make(map[string]string)

	// colors["white"] = "#254278"

	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}