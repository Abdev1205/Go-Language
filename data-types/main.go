package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// creating a new varaible
	var developer string = "Abhay Mishra"
	fmt.Println("Developer: ", developer)

	// short way to create a new varaible
	enginner := "Abhay"
	fmt.Println("Enginner: ", enginner)

	// types of data

	// 1. String
	// deafult value is "" --> empty string
	var stringVar string = "testing"
	stringData := "I am string "
	fmt.Println("StringData: ", stringData, stringVar)

	// we can get len of string by using len fun
	fmt.Println("len of string: ", len(stringData))

	// but above approach we not give correct answer all time because it give wrong answer because char outside the ascii code go use extra bits to store that characters so it is advised to import a unicode/utf8 package and use utf8.RuneCountInString(string)

	fmt.Println(utf8.RuneCountInString(stringData))

	// 2. INT
	// in int we have we have 4 types of types
	//  int8, int16, int32, int64
	// deafult value is 0
	var intVar int64 = 24
	fmt.Println(intVar)

	// 3. Bool
	// deafult value is false
	var boolvar bool = false
	fmt.Println(boolvar)
}
