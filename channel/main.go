package main

import (
	"fmt"
	"time"
)

var (
	arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

func main() {
	fmt.Println("go channel")
	withoutGo()
}

func withGo() {
	defer timer()()
	chaRef1 := make(chan int)
	chaRef2 := make(chan int)

	go func() {
		chaRef1 <- sum(arr1)
	}()

	go func() {
		chaRef2 <- sum(arr2)
	}()

	fmt.Println("sum arr1: ", <-chaRef1)
	fmt.Println("sum arr2: ", <-chaRef2) // cammand + D ใว้ select หลายๆค่าที่เหมืิอนกันในแถวเดียวกัน
}

func withoutGo() {
	defer timer()()
	fmt.Println("sum arr1: ", sum(arr1))
	fmt.Println("sum arr2: ", sum(arr2)) // cammand + D ใว้ select หลายๆค่าที่เหมืิอนกันในแถวเดียวกัน
}

func timer() func() {
	t := time.Now()

	return func() {
		diff := time.Now().Sub(t)
		fmt.Println(diff)
	}
}

func sum(a []int) int {
	s := 0
	for _, x := range a {
		s += x
		time.Sleep(time.Millisecond * 200)
	}
	return s
}
