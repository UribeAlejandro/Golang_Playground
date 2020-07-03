package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	for {
		fmt.Println("Please enter a string.")
		_, err := fmt.Scan(&s)
		s = strings.ToLower(s)

		if strings.Contains(s, "a") == true && strings.HasSuffix(s, "n") == true && strings.HasPrefix(s, "i") == true {
			fmt.Println("Found")

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Not Found")
		}
	}
}
