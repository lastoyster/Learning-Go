package main

import "fmt"

func main() {
	menu := [6]string{"Chocolate", "Cake", "Coffee", "Crab", "Crunchy", "Fries"}

	var x int
	fmt.Scanln(&x)
	if x < 6 {
		fmt.Println(menu[x])
	} else {
		fmt.Println("Invalid choice")
	}
}
