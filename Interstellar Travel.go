//You’ll be writing Go functions to perform calculations and build out an interstellar travel agency! by @codecademy
package main

import (
	"fmt"
)

func fuelGauge(fuel int) {
	fmt.Printf("Current available fule: %d \n", fuel)
}

func calculateFuel(planet string) int {
	var fuel int
	
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}

	return fuel
}

func greetPlanet(planet string) {
	fmt.Printf("Mic check! We've reached your destination! Planet %s \n", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}

func main() {
	fuel := 1000000
	planetChoice := "Venus"

	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)

	planetChoice = "Mars"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
