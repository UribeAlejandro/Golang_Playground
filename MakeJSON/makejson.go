package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
Write a program which prompts the user to first enter a name,
and then enter an address. Your program should create a map and add the
name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and
then your program should print the JSON object.

*/

func main() {
	person := make(map[string]string, 2)
	person["name"] = ""
	person["address"] = ""
	scanner := bufio.NewScanner(os.Stdin)
	for {
		for key, _ := range person {
			fmt.Println("Please enter your", key)
			scanner.Scan()
			person[key] = scanner.Text()
		}
		barr, err := json.Marshal(person)
		if err == nil {
			fmt.Println(string(barr))
			os.Exit(1)
		}
	}
}
