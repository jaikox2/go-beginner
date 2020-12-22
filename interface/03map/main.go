package main

import "fmt"

type person struct {
	name string
}

func main() {
	fmt.Println("======interface with map=======")

	m := make(map[interface{}]interface{})
	m[1] = "Hello"
	m["name"] = "Pang"
	m["x"] = 10

	p := person{
		"Gopher",
	}
	fmt.Println(m)
	m[p] = "Takashi"
	fmt.Println(m[p])

	fmt.Println("======interface with map check type=======")

	fmt.Println(m["y"])
	if x, ok := m["name"].(string); ok {
		fmt.Println(x)
	}

}
