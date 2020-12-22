package main

import "fmt"

func main() {
	var a [2][3]int

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = j
		}
	}

	fmt.Println(a)
}
