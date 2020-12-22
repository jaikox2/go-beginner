package main

import "fmt"

type person struct {
	Name string
}

func (p person) Talk() {
	fmt.Println("Hello, I am ", p.Name)
}

type cat struct{}

func (cat) Talk() {
	fmt.Println("Meo Meo")
}

type dog struct{}

func (*dog) Talk() {
	fmt.Println("Wan Wan")
}

// shot key => tyi
type talkable interface {
	Talk() // auto implement
}

func talkWith(t talkable) {
	t.Talk()
}

func main() {
	fmt.Println("=========Interface=========")
	p := person{
		Name: "Pang",
	}
	p.Talk()

	talkWith(p)
	talkWith(cat{})
	talkWith(&dog{})
}
