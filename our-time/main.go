package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study ")

	currTime := time.Now()
	fmt.Println("Our curr time is ", currTime)

	// we can also use formater to get the value
	// we have .format in go lang to get the value
	// comman pardiagm is used format 01-02-2006 Monday

	fmt.Println("Formated curr time is ", currTime.Format("01-02-2006 Monday"))

	// we can also creata a Date using .Date

	createdData := time.Date(2024, time.December, 12, 21, 232, 34, 232, time.UTC)

	fmt.Println("Created date is ", createdData)
}
