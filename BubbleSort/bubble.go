package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := InputArray()
	BubbleSort(x)
	fmt.Println("Please find below the sorted array.")
	fmt.Println(x)
}

func BubbleSort(x []int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				Swap(x, j)
			}
		}
	}
}

func Swap(x []int, i int) {
	temp := x[i]
	x[i] = x[i+1]
	x[i+1] = temp
}

func InputArray() []int {
	var x []int
	fmt.Println("Please enter a sequence of integers.")
	fmt.Println("Separated by spaces, max 10.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	xt := strings.Split(s, " ")
	for _, val := range xt {
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		} else {
			x = append(x, i)
		}
	}
	x = TruncArray(x)
	return x
}

func TruncArray(x []int) []int {
	fmt.Println()
	fmt.Println("You entered", len(x), "integers.")
	if len(x) == 11 {
		fmt.Println("The last integer was eliminated.")
		x = x[0:10]
	} else if len(x) > 11 {
		fmt.Println("The last", len(x)-10, "integers were eliminated.")
		x = x[0:10]
	}
	fmt.Println()
	return x
}
