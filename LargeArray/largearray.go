package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x []int
	var y []int
	var w []int

	ah := make(chan []int)
	bh := make(chan []int)
	ch := make(chan []int)
	dh := make(chan []int)

	WelcomeMessg()
	x = UserInput()
	a, b, c, d := Splitter(x)

	fmt.Println("Sublists to be sorted.")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	go SortSlice(a, ah)
	go SortSlice(b, bh)
	go SortSlice(c, ch)
	go SortSlice(d, dh)

	a = <-ah
	b = <-bh
	c = <-ch
	d = <-dh

	//Takes two chunks and re-sort
	x = append(a, b...)
	y = append(c, d...)

	go SortSlice(x, ah)
	go SortSlice(y, bh)

	x = <-ah
	y = <-bh

	//Checks sorting of two last chunks together
	w = append(x, y...)

	go SortSlice(w, ah)

	w = <-ah
	fmt.Println("Sorted list:", w)
}

func UserInput() []int {
	var s string
	var x []int
	var xt []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()
	xt = strings.Split(s, " ")
	if len(xt) > 3 {
		for _, val := range xt {
			i, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
				UserInput()
			} else {
				x = append(x, i)
			}
		}
	} else {
		fmt.Println("Please enter four or more numbers.")
		UserInput()
	}
	return x
}

func Splitter(x []int) ([]int, []int, []int, []int) {
	var a []int
	var b []int
	var c []int
	var d []int
	var l int

	l = len(x) / 4

	a = x[0:l]
	b = x[l : 2*l]
	c = x[2*l : 3*l]
	d = x[3*l : len(x)]

	return a, b, c, d
}

func WelcomeMessg() {
	fmt.Println("Please enter a sequence of integers.")
	fmt.Println("Separated by spaces.")
}

func SortSlice(x []int, ch chan []int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x = Swap(x, j)
			}
		}
	}
	ch <- x
}

func Swap(x []int, j int) []int {
	temp := x[j]
	x[j] = x[j+1]
	x[j+1] = temp

	return x
}
