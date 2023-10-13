package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The above code defines a struct type called `menuItem` with a name field and a prices field that is
// a map of string keys to float64 values.
// @property {string} name - The name property is a string that represents the name of the menu item.
// @property prices - The `prices` property is a map that associates a string key with a float64 value.
// It is used to store the prices of the menu item for different variations or options.
type menuItem struct {
	name   string
	prices map[string]float64
}

// `in := bufio.NewReader(os.Stdin)` is creating a new `bufio.Reader` object called `in` that reads
// input from the standard input (`os.Stdin`). This allows the program to read user input from the
// command line.
var in = bufio.NewReader(os.Stdin)

// The function "addItem" prompts the user to enter the name of an item and adds it to the menu.
func Add() {
	fmt.Println("Please enter the name of the item: ")
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
}

// The function "printMenu" prints the menu items and their corresponding sizes and prices.
func Print() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price) // see https://pkg.go.dev/fmt#hdr-Printing
		}
	}
}
