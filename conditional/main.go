package main

import "fmt"

func main() {
	var age int = 19
	voter(age)
}

func voter(age int) {
	if age > 18 && age < 60 {
		fmt.Println("You can vote")
	}
	// elif(age>60){
	// 	fmt.PrintLn("You can vote)
	// }
}
