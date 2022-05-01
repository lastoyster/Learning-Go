package main

import "fmt"

func main() {
	var age int
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars)
}
func mars_age(a int) int {
	var res int
	res = (a * 365)
	return res / 687
}
