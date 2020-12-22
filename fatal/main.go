package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("========fatal========")
	fmt.Println("Start.....")
	log.Fatal("fatal error, program can not run")
	fmt.Println("Hello")
}
