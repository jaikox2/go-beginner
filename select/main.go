package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("======SELECT=======")
	res, err := doWorkWithTime(5 * time.Second)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}

func doVeryLongWork() int {
	time.Sleep(4 * time.Second)
	return 1
}

func doWorkWithTime(d time.Duration) (int, error) {
	cha := make(chan int)

	go func() {
		cha <- doVeryLongWork()
	}()

	select {
	case r := <-cha:
		return r, nil
	case <-time.After(d):
		return 0, fmt.Errorf("timeout")
	}
}
