package main

import "fmt"

func main() {
	for {
		UserInput()
		fmt.Println()
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() string {
	return animal.food
}
func (animal Animal) Move() string {
	return animal.locomotion
}
func (animal Animal) Speak() string {
	return animal.noise
}

func UserInput() {
	var animal string
	var action string
	var an Animal

	fmt.Println("Please enter an animal (cow, bird or snake).")
	fmt.Println("Followed by an action (eat, move or speak).")
	fmt.Println("Both separated by a space.")
	fmt.Print(">")
	fmt.Scan(&animal, &action)
	CheckAnimalAction(animal, action)
	switch animal {
	case "cow":
		an = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		an = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	case "snake":
		an = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	}
	switch action {
	case "eat":
		fmt.Println(an.Eat())
	case "move":
		fmt.Println(an.Move())
	case "speak":
		fmt.Println(an.Speak())
	}
}

func CheckAnimalAction(animal, action string) {
	fmt.Println("")
	if (animal == "cow") || (animal == "bird") || (animal == "snake") {
		if (action == "eat") || (action == "move") || (action == "speak") {
		} else {
			fmt.Println("Please enter a valid action.")
			fmt.Println("")
			UserInput()
		}
	} else {
		fmt.Println("Please enter a valid animal.")
		fmt.Println("")
		UserInput()
	}
}
