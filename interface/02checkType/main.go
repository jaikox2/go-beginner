package main

import "fmt"

func main() {
	fmt.Println("====interface 02======")
	checkType("Pang")
	checkType(10)
	checkType(true)

	fmt.Println("====interface check type 2======")
	checkType1("Pang")
	checkType1(10)
	checkType1(true)

	fmt.Println("====interface check type 3======")
	checkType2("Pang")
	checkType2(10)
	checkType2(true)
}

func checkType(v interface{}) {
	switch p := v.(type) {
	case string:
		fmt.Println("String: Hello", p)
	case int:
		fmt.Println("Int:", p+10)
	case bool:
		fmt.Println("Bool", p)
	}
}

func checkType1(v interface{}) {
	p, ok := v.(string)
	if ok {
		fmt.Println(p)
	} else {
		fmt.Println("v is not string type")
	}
}

func checkType2(v interface{}) {
	p, _ := v.(string)
	fmt.Println(p)
}
