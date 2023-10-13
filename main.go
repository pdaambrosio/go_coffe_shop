package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

// The main function allows the user to interact with a menu by printing it, adding items, or quitting
// the program.
func main() {
	// The above code defines a struct type called `menuItem` with a name field and a prices field that is
	// a map of string keys to float64 values.
	// @property {string} name - The name property is a string that represents the name of the menu item.
	// @property prices - The `prices` property is a map that associates a string key with a float64 value.
	// It is used to store the prices of the menu item for different variations or options.
	type menuItem struct {
		name string
		prices map[string]float64
	}

	// The `menu` variable is a slice of `menuItem` structs. Each `menuItem` struct represents an item on
	// the menu and contains two fields: `name` and `prices`.
	menu := []menuItem{
		{name: "Coffe", prices: map[string]float64{"small": 1.50, "medium": 2.00, "large": 2.50}},
		{name: "Tea", prices: map[string]float64{"small": 1.00, "medium": 1.50, "large": 2.00}},
		{name: "Water", prices: map[string]float64{"small": 0.50, "medium": 1.00, "large": 1.25}},
	}

	// `in := bufio.NewReader(os.Stdin)` is creating a new `bufio.Reader` object called `in` that reads
	// input from the standard input (`os.Stdin`). This allows the program to read user input from the
	// command line.
	in := bufio.NewReader(os.Stdin)

// The `loop` label is used to break out of the outer for loop when the user chooses to quit the program.
loop:
	// The for loop is used to repeatedly prompt the user for a menu option until they choose to quit.
	for {
		fmt.Println("Please select an option: ")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add Item")
		fmt.Println("q) Quit")
		choice, _ := in.ReadString('\n')

		// The code you provided is a switch statement that handles the user's menu choice.
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
