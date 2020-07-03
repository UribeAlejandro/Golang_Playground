package main

import (
	"fmt"
	"math"
)

var a float64
var vo float64
var so float64
var t float64

func main() {
	UserInput()
	fn := GenDisplaceFn(a, vo, so)
	t = TakeTime()
	fmt.Println(fn(t))
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	disp_t := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + vo*t + so
	}
	return disp_t
}

func UserInput() {
	fmt.Println("Please enter aceleration, initial velocity and initial displacement. Press enter after each value.")
	_, err := fmt.Scan(&a, &vo, &so)
	if err != nil {
		fmt.Println(err)
		UserInput()
	}
}

func TakeTime() float64 {
	fmt.Println("Please enter a time.")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Println(err)
		TakeTime()
	}
	if t < 0 {
		fmt.Println("Please enter a postive number.")
		TakeTime()
	}

	return t
}
