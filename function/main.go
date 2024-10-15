package main

import (
	"errors"
	"fmt"
)

func main() {
	var yourJob string = "Coding"
	curretntJob(yourJob)

	var numerator int = 10
	var denominator int = 0
	//calling prameter and return type function
	var result, remainder, err = division(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The result of divison is %v and the remainder is %v", result, remainder)
}

// this is no return type funcation example
func curretntJob(yourJob string) {
	fmt.Println(yourJob)
}

// this is return type funcation
// after func() <return_type> is given in funcation and we can also pass multiple return types and return values

func division(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide By Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
