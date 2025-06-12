package main

import "fmt"

func main() {
	fmt.Println("Welcom to map ")

	// creating a map
	language := make(map[string]string)

	language["js"] = "javascript"
	language["go"] = "golang"
	language["rb"] = "ruby"
	language["py"] = "python"

	fmt.Println("language is: ", language)

	// getting the specific language
	fmt.Println("The language of js is: ", language["js"])

	// for deleting we can use the delete keyword
	// this do same as append anithung type

	delete(language, "rb") // so here we are deleting the rb from the map
	fmt.Println("After deleting rb: ", language)

	// we can use also get the data using loop

	for key, val := range language {
		fmt.Println("key: ", key, " value: ", val)
	}
}
