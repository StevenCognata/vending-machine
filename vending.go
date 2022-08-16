package main

import "fmt"

func main() {

	// storing beverage strings in an array
	var beverages = [10]string{"Dr.Pepper", "Coke", "Iced Tea", "Sprite", "Lemonade", "Water", "RedBull", "Orange Soda", "Iced Coffee", "Seltzer"}

	// printing array items to console
	for i := 0; i < len(beverages); i++ {
		fmt.Println(beverages[i])
	}
}
