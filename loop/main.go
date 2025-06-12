package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Go lang")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	// for looping in go lang we have for keyword
	// as in other languages we have while do while but in go
	// we have diff variation in go lang to do same thing here
	// wo avoid confusion and complexity

	// 1st way -> Normal i,j loop with itetation

	for d := 0; d < len(days); d++ { // here we don't have ++d clean and less confusing
		fmt.Println(days[d])
	}

	// 2nd way -> using range here range ieretae thourgh any iteratable data structure

	fmt.Println("\nDays using range")
	for i := range days {
		fmt.Println(days[i])
	}

	// 3rd way -> using for each loop
	// in for each first value is index and second is value
	fmt.Println("\nDays using for each")

	for index, day := range days {
		fmt.Println("Index: ", index, " Day: ", day)
	}

	// Almost loop is cleared

}
