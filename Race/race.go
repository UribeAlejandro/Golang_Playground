package main

import (
	"fmt"
	"time"
)

var x int

func main() {

	for i := 0; i < 10; i++ {
		go inc()
		go printX(i)
	}
	time.Sleep(100 * time.Millisecond)
}

func inc() {

	x++
}

func printX(i int) {
	fmt.Println("Value on iteration", i, ":", x)
}
