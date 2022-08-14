package main

import "fmt"

func main() {

	var beverages = [10]string{"Dr.Pepper", "Coke", "Iced Tea", "Sprite", "Lemonade", "Water", "RedBull", "Orange Soda", "Iced Coffee", "Seltzer"}

	for i := 0; i < len(beverages); i++ {
		fmt.Println(beverages[i])
	}
}
