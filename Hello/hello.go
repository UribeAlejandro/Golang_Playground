package main

import "fmt"

func main() {
	s := "Hello World"
	go Printer(s)
	fmt.Println("nope|")

}
func Printer(s string) {
	fmt.Println(s)
}
