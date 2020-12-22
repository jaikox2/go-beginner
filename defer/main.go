package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=====defer=====")
	f, err := os.Create("file.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() // ถ้า main run เสร็จแล้วให้ปิดด้วย f.Close()
	f.WriteString("Hello")
}
