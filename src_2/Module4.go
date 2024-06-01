package main

import (
	"fmt"
)

// Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow struct
type Cow struct{}

// Eat method for Cow
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Move method for Cow
func (c Cow) Move() {
	fmt.Println("walk")
}

// Speak method for Cow
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird struct
type Bird struct{}

// Eat method for Bird
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Move method for Bird
func (b Bird) Move() {
	fmt.Println("fly")
}

// Speak method for Bird
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake struct
type Snake struct{}

// Eat method for Snake
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Move method for Snake
func (s Snake) Move() {
	fmt.Println("slither")
}

// Speak method for Snake
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	for {
		fmt.Print("> ")
		var command, name, action string
		fmt.Scan(&command, &name, &action)

		switch command {
		case "newanimal":
			var newAnimal Animal
			switch action {
			case "cow":
				newAnimal = Cow{}
			case "bird":
				newAnimal = Bird{}
			case "snake":
				newAnimal = Snake{}
			default:
				fmt.Println("Invalid animal type.")
				continue
			}
			animals[name] = newAnimal
			fmt.Println("Created it!")
		case "query":
			animal, ok := animals[name]
			if !ok {
				fmt.Println("Animal not found.")
				continue
			}
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid action.")
			}
		default:
			fmt.Println("Invalid command.")
		}
	}
}
