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

type menu []menuItem

// Method "print" contains the logic to print the menu.
func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price) // see https://pkg.go.dev/fmt#hdr-Printing
		}
	}
}

// Method "add" contains the logic to add a new item to the menu.
func (m *menu) add() {
	fmt.Println("Please enter the name of the item: ")
	name, _ := in.ReadString('\n')
	data = append(*m, menuItem{name: name, prices: make(map[string]float64)})
}

var in = bufio.NewReader(os.Stdin) // see https://pkg.go.dev/bufio#NewReader

// Function "add" adds a new item to the menu.
func Add() {
	data.add()
}

// Function "Print" prints the menu.
func Print() {
	data.print()
}
