package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// for taking input we have bufio package that works with os
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("PLease Enter the rating ")

	// this , pattern is also called as ok comma or error comma patter
	// varaible can get either value or error
	// there can be ok or any issues | error

	rating, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You entered: %s", rating)
}
