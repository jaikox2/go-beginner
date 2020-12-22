package main

// pkgm = create package main

import "fmt"

// ff = fmt.Printf()
// fp = fmt.Println()

func main() {
	// three of delaration variable in golang
	var gopher string
	fmt.Println("gopher: ", gopher)
	gopher = "test"
	fmt.Println("gopher: ", gopher)

	var name = "Irfan kuechi"
	fmt.Printf("My name is %s\n", name)

	nickName := "Pang dev"
	// fmt.Printf("My nick name is %s\n", nickName)

	// comment variable is not use it
	_ = nickName
}
