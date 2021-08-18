package main

// https://www.codecademy.com/courses/learn-go/projects/learn-go-functions-interstellar-travel

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Printf("There is %d gallons of fuel left\n", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mecury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Printf("Welcome to %v\n", planet)
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func willYouMakeIt(planet string, fuel int) bool {
	var fuelRemaining, fuelCost int
	var success bool
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		success = true
	}
	if fuelCost > fuelRemaining {
		cantFly()
		success = false
	}

	return success
}

// Create the function flyToPlanet() here
func flyToPlanet(startLocation, targetPlanet string, fuel int) (int, string) {
	var fuelRemaining, fuelCost int
	var currentLocation string

	fuelRemaining = fuel
	currentLocation = startLocation
	fuelCost = calculateFuel(targetPlanet)

	if willYouMakeIt(targetPlanet, fuel) {
		greetPlanet(targetPlanet)
		fuelRemaining -= fuelCost
		currentLocation = targetPlanet
	}

	return fuelRemaining, currentLocation
}

func whereAreWe(currentLocation string) {
	fmt.Printf("We are on %v\n\n", currentLocation)
}

func travel(currentLocation, planetChoice string, fuel int) (int, string) {
	fuel, currentLocation = flyToPlanet(currentLocation, planetChoice, fuel)
	fuelGauge(fuel)
	whereAreWe(currentLocation)

	return fuel, currentLocation
}

func main() {
	// Create `planetChoice` and `fuel`
	var fuel int
	var planetChoice string
	var currentLocation string

	fuel = 1000000
	currentLocation = "Earth"
	planetChoice = "Venus"

	fuel, currentLocation = travel(currentLocation, planetChoice, fuel)
	fuel, currentLocation = travel(currentLocation, "Mars", fuel)
	fuel, currentLocation = travel(currentLocation, "Mecury", fuel)

}
