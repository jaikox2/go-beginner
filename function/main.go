package main

import "fmt"

func main() {
	fmt.Println("a")
	a, b := 1, 2
	fmt.Println(add(a, b))

	fmt.Println(add(4, 4))

	fmt.Println("=====array=====")

	x := [5]int{}
	fmt.Println(x)
	mutateArray(x)
	fmt.Println(x)

	fmt.Println("=====slice=====")
	mutateSlice(x[0:len(x)]) // change array to slice
	fmt.Println(x)
	mutateSlice(x[:]) // shot
	fmt.Println(x)
	mutateSlice(x[2:]) // shot
	fmt.Println(x)
}

func add(x, y int) int {
	return x + y
}

func mutateArray(a [5]int) {
	a[0] = 10
	fmt.Println(a)
}

func mutateSlice(a []int) {
	a[0] = 10
}
