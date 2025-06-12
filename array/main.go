package main

import "fmt"

func main() {
	// we can create a array in this way
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"

	fmt.Println("Fruits without direct intialization ", fruits)

	// we can also create and initialize also in the same way
	var veg = [3]string{"bhindi", "Kaddu", "Gobhi"}
	fmt.Println("Vegetables ", veg)

	// we can also get the array len using len(fruits)

	fmt.Println("Len of fruits ", len(fruits))

}
