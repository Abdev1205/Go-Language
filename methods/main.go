package main

import "fmt"

type Student struct {
	Name      string
	Age       int
	Grade     int
	isPresent bool
}

// So Now I will create a new methods that will tell student is present or not
// so here we will add struct as copy object as same in c++
func (s Student) getStatus() bool {
	return s.isPresent
}

func (s Student) fullName() {
	s.Name = "Abhay Mishra"
	fmt.Println("Full Name: ", s.Name)
}

func main() {
	fmt.Println("Welcome to Methods in Go lang")

	// Methods is same as funcation but it applies for class
	// so in go lang we can have thing with struct
	// we can define the function as expect param as struct
	// ex somethods(s someStructs) // can we can use s uinsg dot operator

	abhay := Student{Name: "Abhay", Age: 22, Grade: 10, isPresent: true}
	fmt.Println("Abhay is present: ", abhay.getStatus())
	fmt.Println("Abhay full name: ")
	abhay.fullName()

	// but as we print again the same thing we see that Struct is passed as copy
	// so it will not change the object
	// to avoid this we have to pass as the refernce using pointer
}
