package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x string

	for {
		fmt.Printf("Please enter a float.")

		_, err := fmt.Scan(&x)

		xt, err := strconv.ParseFloat(x, 64)

		fmt.Println(int64(xt))

		if err != nil {
			fmt.Println(err)
		}
	}
}
