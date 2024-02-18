package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var cow animal
	var bird animal
	var snake animal

	cow.initAnimal("cow", "grass", "walk", "moo")
	bird.initAnimal("bird", "worms", "fly", "peep")
	snake.initAnimal("snake", "mice", "slither", "hiss")

	fmt.Println("Welcome to the Animal Information Program!")
	fmt.Println("Enter 'quit' to exit.")

	for {
		fmt.Print("> ")

		// Read user input
		reader := bufio.NewReader(os.Stdin)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Trim whitespace and split the input
		userInput = strings.TrimSpace(userInput)
		inputSlice := strings.Split(userInput, " ")

		// Check if the input is valid
		if len(inputSlice) != 2 {
			fmt.Println("Invalid input! Please enter two strings separated by space.")
			continue
		}

		// Check if the action is valid
		name := inputSlice[0]
		if name != "cow" && name != "bird" && name != "snake" {
			fmt.Println("Invalid animal! Please enter 'cow', 'bird', or 'snake'.")
			continue
		}

		// Check if the action is valid
		action := inputSlice[1]
		if action != "eat" && action != "move" && action != "speak" {
			fmt.Println("Invalid action! Please enter 'eat', 'move', or 'speak'.")
			continue
		}

		//fmt.Println(name)
		//fmt.Println(action)

		if name == "cow" && action == "move" {
			cow.move()
		} else if name == "cow" && action == "eat" {
			cow.eat()
		} else if name == "cow" && action == "speak" {
			cow.speak()
		} else if name == "bird" && action == "move" {
			bird.move()
		} else if name == "bird" && action == "eat" {
			bird.eat()
		} else if name == "bird" && action == "speak" {
			bird.speak()
		} else if name == "snake" && action == "move" {
			snake.move()
		} else if name == "snake" && action == "eat" {
			snake.eat()
		} else if name == "snake" && action == "speak" {
			snake.speak()
		}
	}

}

type animal struct {
	animalName  string
	foodEaten   string
	locomotion  string
	spokenSound string
}

func (a *animal) printAnimalName() {
	fmt.Println(a.animalName)
}
func (a *animal) eat() {
	fmt.Println(a.foodEaten)
}
func (a *animal) move() {
	fmt.Println(a.locomotion)
}
func (a *animal) speak() {
	fmt.Println(a.spokenSound)
}

func (a *animal) initAnimal(animalName string, foodEaten string, locomotion string, spokenSound string) {
	a.animalName = animalName
	a.foodEaten = foodEaten
	a.locomotion = locomotion
	a.spokenSound = spokenSound
}
