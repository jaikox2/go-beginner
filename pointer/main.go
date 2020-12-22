package main

import "fmt"

func main() {
	fmt.Println("Pointer")
	a := 10
	fmt.Println(a)

	ptrA := &a
	fmt.Println(ptrA)
	fmt.Println(*ptrA)

	*ptrA = 20
	fmt.Println(a)

}
