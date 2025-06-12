package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defer in Go lang")

	// defer is used to when we want to have the certain line to execute at the last
	// This is lifo which mean last in first out so the last defer will be executed first

	defer fmt.Println("Hello")
	fmt.Println("World")
}
