package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"coffeShop/menu"
)

// `in := bufio.NewReader(os.Stdin)` is creating a new `bufio.Reader` object called `in` that reads
// input from the standard input (`os.Stdin`). This allows the program to read user input from the
// command line.
var in = bufio.NewReader(os.Stdin)

func main() {
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
