package main

import (
	"fmt"
)

func main() {
	fmt.Println("========panic========")
	fmt.Println("Start.....")
	// panic("panic error, program can not run")
	doSafework()
	fmt.Println("Done")
}

func doFailWork() {
	panic("panic fail")
}

func doSafework() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("work fail: ", r)
		}
	}()
	doFailWork()
	fmt.Println("work success")
}
