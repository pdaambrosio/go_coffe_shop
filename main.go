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

// The main function allows the user to interact with a menu by printing it, adding items to it, or
// quitting the program.
func main() {
	type menuItem struct {
		name string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffe", prices: map[string]float64{"small": 1.50, "medium": 2.00, "large": 2.50}},
		{name: "Tea", prices: map[string]float64{"small": 1.00, "medium": 1.50, "large": 2.00}},
		{name: "Water", prices: map[string]float64{"small": 0.50, "medium": 1.00, "large": 1.25}},
	}

	in := bufio.NewReader(os.Stdin)

loop:
	for {
		fmt.Println("Please select an option: ")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add Item")
		fmt.Println("q) Quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.prices {
					fmt.Printf("\t%10s%10.2f\n", size, price) // see https://pkg.go.dev/fmt#hdr-Printing
				}
			}
		case "2":
			fmt.Println("Please enter the name of the item: ")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name:name, prices:make(map[string]float64)})
		case "q":
			fmt.Println("Goodbye!")
			break loop
		default:
			fmt.Println("Invalid option")
		}
	}
}
