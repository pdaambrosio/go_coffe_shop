package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"coffeShop/menu"
)

var in = bufio.NewReader(os.Stdin) // see https://pkg.go.dev/bufio#NewReader

// Function "main" is the entry point of the program.
func main() {
loop:
	for {
		fmt.Println("Please select an option: ")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add Item")
		fmt.Println("q) Quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			menu.Add()
		case "q":
			fmt.Println("Goodbye!")
			break loop
		default:
			fmt.Println("Invalid option")
		}
	}
}
