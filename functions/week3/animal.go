package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func getRequest() (string, string) {
	var animalChoice, attribute string
	for {
		fmt.Print(">")
		_, error := fmt.Scanln(&animalChoice, &attribute)
		if error != nil {
			fmt.Println("Error", error)
		} else {
			break
		}
	}
	return animalChoice, attribute
}

func main() {
	fmt.Println("pls enter '[cow|bird|snake] [eat|move|speak]'")
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	printDetails(cow, bird, snake)
}

func printDetails(cow Animal, bird Animal, snake Animal) {
	for {
		selectedAnimal, attibute := getRequest()

		var animal Animal
		if selectedAnimal == "cow" {
			animal = cow
		} else if selectedAnimal == "bird" {
			animal = bird
		} else if selectedAnimal == "snake" {
			animal = snake
		}
		if attibute == "eat" {
			animal.Eat()
		} else if attibute == "move" {
			animal.Move()
		} else if attibute == "speak" {
			animal.Speak()
		}
	}
}
