package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Switch in Go ")
	dice := rand.Intn(6) + 1
	fmt.Println("Dice Value: ", dice)

	switch dice {
	case 1:
		fmt.Println("You got a 1")
	case 2:
		fmt.Println("You got a 2")
	case 3:
		fmt.Println("You got a 3")
	case 4:
		fmt.Println("You got a 4")
	case 5:
		fmt.Println("You got a 5")
	case 6:
		fmt.Println("You got a 6")
	default:
		fmt.Println("Invalid Dice Value")
	}
}
