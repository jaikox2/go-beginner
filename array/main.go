package main

import "fmt"

func main() {
	// var a [5]int // ค่าเริ่มต้นของตัวแปร int = 0, string = "", bool = false

	// a[0] = 10
	// a[2] = 30
	// a[4] = 40

	a := [5]int{2, 4, 5}

	fmt.Println(a)
	fmt.Println("number of array :", len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println("another for ->\n")

	for key, value := range a {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}

	fmt.Println("another for ->")

	for _, value := range a {
		// fmt.Println(key)
		fmt.Println(value)
	}

	// shot key => forr or for

	var b bool

	fmt.Println(b)
}
