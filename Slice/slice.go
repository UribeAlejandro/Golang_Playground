package main

import (
	"fmt"
	"sort"
)

func main() {
	s := make([]int, 0, 3)
	var x int
	i := 0
	for {
		fmt.Println("Please enter an integer.")
		_, err := fmt.Scan(&x)

		if err == nil {
			s = append(s, x)
			sort.Ints(s[0 : i+1])
			fmt.Println(s[0 : i+1])
			fmt.Println()
			i++
		} else {
			fmt.Print("Error: ")
			fmt.Println(err)
		}
	}
}
