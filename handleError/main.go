package main

import (
	"errors"
	"fmt"
)

var (
	ageTooLow  = errors.New("age too low")
	ageTooHigh = errors.New("age too low")
)

func validateAge(age int) error {
	if age < 18 {
		// return fmt.Errorf("age too low") // can bind value like Printf
		return ageTooLow
	} else if age > 60 {
		// return fmt.Errorf("age too high")
		return ageTooHigh
	}
	return nil
}

func main() {
	fmt.Println("=======Handle Error======")
	fmt.Println(validateAge(50))

	err := validateAge(70)

	if err == ageTooLow {
		fmt.Println("CAN NOT ENTER")
		return
	}
	if err == ageTooHigh {
		fmt.Println(":D")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
