// Package "menu" implements the menu for the coffee shop.
package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Struct "menuItem" are the items on the menu. Each item has a name and a map of prices for each size.
type menuItem struct {
	name   string
	prices map[string]float64
}

var in = bufio.NewReader(os.Stdin) // see https://pkg.go.dev/bufio#NewReader

// Function "add" adds a new item to the menu.
func Add() {
	fmt.Println("Please enter the name of the item: ")
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
}

// Function "Print" prints the menu.
func Print() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price) // see https://pkg.go.dev/fmt#hdr-Printing
		}
	}
}
