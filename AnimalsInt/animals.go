package main

import (
	"fmt"
	"strings"
)

//Initializes variables
var zoo []Animal
var an Animal

func main() {
	WelcomeMessage()
	UserRequest()
}

type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

//Initializes Cow and its methods
type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}
func (c Cow) GetName() string {
	return c.name
}

//Initializes Bird and its methods
type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}
func (b Bird) GetName() string {
	return b.name
}

//Initializes Snake and its methods
type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}
func (s Snake) GetName() string {
	return s.name
}

func UserRequest() {
	var req string
	var name string
	var acttyp string

	fmt.Print(">")
	for {
		_, err := fmt.Scan(&req, &name, &acttyp)
		req = strings.ToLower(req)
		acttyp = strings.ToLower(acttyp)
		if err == nil {
			if req == "newanimal" {
				NewAnimal(name, acttyp)
			} else if req == "query" {
				Query(name, acttyp)
			} else {
				fmt.Println("Wrong request:", req)
				UserRequest()
			}
		} else {
			fmt.Println(err)
		}
	}
}

func NewAnimal(name, acttyp string) {
	switch acttyp {
	default:
		fmt.Println("Wrong type:", acttyp)
		fmt.Println("")
		UserRequest()
	case "cow":
		an = Cow{name}
	case "bird":
		an = Bird{name}
	case "snake":
		an = Snake{name}
	}
	fmt.Println("Created it!")
	zoo = append(zoo, an)
	UserRequest()
}

func Query(name, acttyp string) {
	var createA string
	if len(zoo) > 0 {
		for i, _ := range zoo {
			if zoo[i].GetName() == name {
				switch acttyp {
				case "eat":
					zoo[i].Eat()
				case "move":
					zoo[i].Move()
				case "speak":
					zoo[i].Speak()
				default:
					fmt.Println("Wrong action:", acttyp)
				}
				UserRequest()
			}
		}
	} else {
		fmt.Println(name, "was not found, do you want to create it? [y/n]")
		_, err := fmt.Scan(&createA)
		createA = strings.ToLower(createA)
		if err == nil {
			if createA == "y" {
				fmt.Println("Please enter its type:")
				fmt.Print("newanimal ", name, " ")
				_, err := fmt.Scan(&acttyp)
				if err == nil {
					NewAnimal(name, acttyp)
				} else {
					fmt.Println(err)
				}
			} else {
				UserRequest()
			}
		} else {
			fmt.Println(err)
		}
	}

}

func WelcomeMessage() {
	fmt.Println("Please enter your request, either newanimal or query.")
	fmt.Println("A newanimal request should be followed by its name and its type: cow, bird or snake.")
	fmt.Println("A query request should be followed by an animal's name and an action: eat, move or speak.")
}
