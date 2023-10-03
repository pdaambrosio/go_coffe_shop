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

func main() {
fmt.Println("Please select an option: ")
fmt.Println("1) Print Menu")

in := bufio.NewReader(os.Stdin)
choice, _ := in.ReadString('\n')
choice = strings.TrimSpace(choice)

// The above type represents a menu item with a name and prices for different options.
// @property {string} name - The name property is a string that represents the name of the menu item.
// @property prices - The "prices" property is a map that associates a string key with a float64 value.
// It is used to store the prices of the menu item for different variations or options. The string key
// represents the variation or option, and the float64 value represents the price for that variation or
// option.
type menuItem struct {
	name string
	prices map[string]float64
}

// The code `menu := []menuItem{...}` is creating a slice of `menuItem` structs. Each `menuItem` struct
// represents a menu item with a name and a map of prices for different sizes. The slice is initialized
// with three `menuItem` structs representing "Coffee", "Tea", and "Water" menu items, along with their
// respective prices for small, medium, and large sizes.
menu := []menuItem{
	{name: "Coffe", prices: map[string]float64{"small": 1.50, "medium": 2.00, "large": 2.50}},
	{name: "Tea", prices: map[string]float64{"small": 1.00, "medium": 1.50, "large": 2.00}},
	{name: "Water", prices: map[string]float64{"small": 0.00, "medium": 0.00, "large": 0.00}},
}

// TODO - Other menu options and logic will go here in future lessons.

fmt.Println(menu)
}
