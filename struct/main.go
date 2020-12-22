package main

import "fmt"

// shot key => tys
type person struct {
	Name     string
	Nickname string
}

// method of struct
func (p *person) mutate() {
	p.Name = "Method struct"
	p.Nickname = "Method"
}

// function
func mutatePerson(p *person) {
	p.Name = "Hacker"
	p.Nickname = "Hack"
}

func main() {
	fmt.Println("=====Struct======")

	p1 := person{
		Name:     "Pang",
		Nickname: "Pang dev",
	}

	fmt.Println(p1)

	fmt.Println("=====P2======")

	p2 := person{
		"Pang",
		"Pang dev",
	}

	fmt.Println(p2)

	fmt.Println("=====P3======")

	p3 := struct {
		Name     string
		Nickname string
	}{
		"Default",
		"Default dev",
	}

	fmt.Println(p3)

	fmt.Println("=====P4======")

	var p4 person

	p4.Name = "Jone"
	p4.Nickname = "Jone dev"

	fmt.Println(p4)

	fmt.Println("=====P4 function======")

	// function
	mutatePerson(&p4)
	fmt.Println(p4)

	fmt.Println("=====P4 method of struct======")

	// method
	p4.mutate()
	fmt.Println(p4)
	fmt.Println(p4.Name)

}
