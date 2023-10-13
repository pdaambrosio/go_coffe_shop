// Package "menu" implements the menu for the coffee shop.
package menu

// Variable "menu" is a slice of "menuItem" structs.
var menu = []menuItem{
	{name: "Coffe", prices: map[string]float64{"small": 1.50, "medium": 2.00, "large": 2.50}},
	{name: "Tea", prices: map[string]float64{"small": 1.00, "medium": 1.50, "large": 2.00}},
	{name: "Water", prices: map[string]float64{"small": 0.50, "medium": 1.00, "large": 1.25}},
}
