package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices are going ont ")
	/*
		Slices are same as arary but here we don;t have specify the size of array
		here we have to intialize with the value
	*/

	var fruites = []string{"apple,mango"}
	fmt.Printf("Type of the fruits is %T \n", fruites)
	// fmt.Println(fruites)

	/*
		we can add data in slice using appned methods
		which take the slice to bemodified amd data
		multiple data can added by adding ,
	*/

	fruites = append(fruites, "Papaya", "Banana")
	fmt.Println("New fruit", fruites)

	/*
		We can also use [:] syntax to slice data in slice
	*/

	fruites = append(fruites[1:]) // this will exlude the firts vlue

	fmt.Println("After colon ", fruites)

	/*
		There is another way to create  a slice or array using make
		here make need to specify the type and size
	*/

	veg := make([]int, 4)

	veg[2] = 10
	veg[1] = 20
	veg[1] = 30
	veg[3] = 40

	fmt.Println("Veg", veg)

	// we can also sort the data useing sort method and it also have various sub methods

	sort.Ints(veg)
	fmt.Println("Sorted veg", veg)

	// wee can check that slice or array is sorted or not

	isSorted := sort.IntsAreSorted(veg)
	fmt.Println("Is sorted", isSorted)

}
