package main

import "fmt"

func main() {
	// a := make([]int, 5)

	// a[2] = 2
	// a[3] = 4

	// a = append(a, 40)

	a := []int{1, 2, 3, 4, 5, 6}

	b := a[:3]

	b[0] = 40         // make b be a part of a // b changed make a change too
	b = append(b, 30) // make it append on a[2]

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a[2:4]) // 2 -> (4-1)
	fmt.Println(a[0:4]) // 0-> (4-1)
	fmt.Println(a[:4])  // 0 -> (4-1)
	fmt.Println(a[4:])  // 4 -> ...
}
