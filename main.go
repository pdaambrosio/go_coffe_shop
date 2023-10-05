package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

/*
This is a simple program that will allow the user to select a menu option such as:
1) Print Menu
*/

// The main function prints a menu of items with their corresponding prices.
// The user can select an option to print the menu.
func main() {
	fmt.Println("Please select an option: ")
	fmt.Println("1) Print Menu")

	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)

	type menuItem struct {
		name string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffe", prices: map[string]float64{"small": 1.50, "medium": 2.00, "large": 2.50}},
		{name: "Tea", prices: map[string]float64{"small": 1.00, "medium": 1.50, "large": 2.00}},
		{name: "Water", prices: map[string]float64{"small": 0.50, "medium": 1.00, "large": 1.25}},
	}

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price) // see https://pkg.go.dev/fmt#hdr-Printing
		}
	}
}
