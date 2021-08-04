/*
Based on the Bank Hiest CodeAcademy Project
https://www.codecademy.com/courses/learn-go/projects/bank-heist
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var isHeistOn bool = true
	var eludedGuards int = rand.Intn(100) // chance of eluding the guards
	var gangSize int = 2 + rand.Intn(3)

	if eludedGuards < 50 {
		fmt.Println("We made it past the guards")
	} else {
		isHeistOn = false
		fmt.Println("Failed at the first hurdle, the guards stopped us!")
	}

	var opendVault int = rand.Intn(100) // chance of opening the vault

	if isHeistOn && opendVault > 70 {
		fmt.Println("Grab the money and go")
	} else if isHeistOn && opendVault < 70 {
		isHeistOn = false
		fmt.Println("uh oh! The vault is locked!")
	}

	var leftSafely int = rand.Intn(5) // chance of making out the bank safely

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Rats the guards were waiting outside the vault...")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("There was a hole in the swag bag; you dropped all the money...")
		case 3:
			isHeistOn = false
			fmt.Println("A brave citizen tripped you up on your way out, giving the guards enough time to grab you...")
		case 4:
			fmt.Println("Jump in the car and floor it!")
		}
	}

	if isHeistOn {
		var amtStolen int = 10000 + rand.Intn(1000000)
		fmt.Printf("Good Work gang! We got away with £%v\n", amtStolen)
		fmt.Printf("We get £%v split between the %v of us!\n", amtStolen/gangSize, gangSize)
	}
}
