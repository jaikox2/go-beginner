package main

import "fmt"

func main() {
	fmt.Print("Input name of a fruit: ")
	var fruit string
	fmt.Scanln(&fruit)

	if fruit == "" {
		fmt.Println("meh! ")
		// return
	}

	if len(fruit) == 0 {
		fmt.Println("meh! ")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("Apple")
	case "banana":
		fmt.Println("Banana")
	case "orange":
		fmt.Println("Orange")
	default:
		fmt.Println("Default")

	}
}
