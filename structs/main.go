package main

import "fmt"

type User struct {
	Name      string
	Age       int64
	Email     string
	IsStudent bool
}

func main() {
	fmt.Println("Welcom to Structs")

	// As in go lang we dont have class so we have struct but it do not have inheritance
	// struct name and its varibael shoudl be in capital
	// similar like we are having data simialar to passing data to constructor

	abhay := User{"Abhay", 22, "abhaym1205@gmail.com", true}
	fmt.Println("Abhay User: ", abhay)

	// if we want to have more details value so we can used %+v
	fmt.Printf("User with exta info %+v \n", abhay) // this will include keys and values good for debug

	// we can use dot operator to get the value

	fmt.Println("Name: ", abhay.Name)

}
