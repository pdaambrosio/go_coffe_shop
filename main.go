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
// TODO: Create a function that will print the menu
}
