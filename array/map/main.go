package main

import "fmt"

func main() {
	a := make(map[string]string)

	a["hello"] = "Gopher"
	a["name"] = "Pang dev"

	fmt.Println(a)
	fmt.Println(a["hello"])
	fmt.Println(a["name"])

	fmt.Println(a["x"])

	fmt.Println("========test 1==========")

	x, ok := a["x"]
	fmt.Println(ok)
	fmt.Println(x)

	fmt.Println("==========test 2==========")

	a["x"] = ""
	x3, ok3 := a["x"]
	fmt.Println(ok3)
	fmt.Println(x3)

	fmt.Println("==========test 3==========")

	a["x"] = "Y"
	x2, ok2 := a["x"]
	fmt.Println(ok2)
	fmt.Println(x2)

	fmt.Println("==========test 4==========")

	delete(a, "x")
	x4, ok4 := a["x"]
	fmt.Println(ok4)
	fmt.Println(x4)

	fmt.Println("==========test 5==========")

	if x5, ok5 := a["x"]; ok5 {
		fmt.Println(x5)
	}

	a["x"] = "TT"

	if x5, ok5 := a["x"]; ok5 {
		fmt.Println(x5)
	}

	fmt.Println("==========test 6==========")

	for key, value := range a {
		fmt.Println("key:", key, "\nvalue: ", value)
	}

	fmt.Println("==========test 7==========")

	b := map[string]string{
		"hello": "GO",
		"name":  "TATA",
	}

	fmt.Println(b)
	fmt.Println(b["name"])

}
