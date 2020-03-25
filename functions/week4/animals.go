package main

import (
	"errors"
	"fmt"
)

type Animal interface {
	getName() string
	Eat()
	Move()
	Speak()
}

// define the cow struct and override the methods in the Animal interface
type Cow struct{ name string }

func (cow Cow) getName() string { return cow.name }
func (cow Cow) Eat()            { fmt.Println("grass") }
func (cow Cow) Move()           { fmt.Println("walk") }
func (cow Cow) Speak()          { fmt.Println("moo") }

// define the bird struct and override the methods in the Animal interface
type Bird struct{ name string }

func (bird Bird) getName() string { return bird.name }
func (bird Bird) Eat()            { fmt.Println("worms") }
func (bird Bird) Move()           { fmt.Println("fly") }
func (bird Bird) Speak()          { fmt.Println("peep") }

// define the snake struct and override the methods in the Animal interface
type Snake struct{ name string }

func (snake Snake) getName() string { return snake.name }
func (snake Snake) Eat()            { fmt.Println("mice") }
func (snake Snake) Move()           { fmt.Println("slither") }
func (snake Snake) Speak()          { fmt.Println("hsss") }

// validate commands
func validateCommand(command string) bool {
	return command == "newanimal" || command == "query"
}

// validate animals
func validateAnimal(animal string) bool {
	return animal == "cow" || animal == "bird" || animal == "snake"
}

// validate annimal prop
func validateProperty(property string) bool {
	return property == "eat" || property == "move" || property == "speak"
}

func getUserInput() (string, string, string) {
	var command, name, option string
	for {
		fmt.Print(">")
		if _, err := fmt.Scanln(&command, &name, &option); err != nil {
			fmt.Println("Error:", err)
		} else {
			if validateCommand(command) {
				if command == "newanimal" {
					if validateAnimal(option) {
						break
					} else {
						fmt.Println("Invalid animal. valid [cow, bird, snake]")
					}
				} else if command == "query" {
					if validateProperty(option) {
						break
					} else {
						fmt.Println("Invalid animal prop. valid [eat, move, speak]")
					}
				}
			} else {
				fmt.Println("Invalid command. [newanimal, query].")
			}
		}
	}
	return command, name, option
}

func appendToAnimalArr(animals *[]Animal, animal Animal) {
	*animals = append(*animals, animal)
	fmt.Println("Created it!")
}

func findAnimalByName(animals []Animal, name string) (Animal, error) {
	var foundAnimal Animal
	var err = errors.New("animal not found !")
	for _, animal := range animals {
		if animal.getName() == name {
			foundAnimal = animal
			err = nil
			break
		}
	}
	return foundAnimal, err
}

/*
 * how to test this ?
 * newanimal jess cow
 * query jess eat
 */
func main() {
	// will be used to store animals
	var animalArr []Animal
	for {
		// read user input
		command, name, option := getUserInput()
		switch command {
		case "newanimal":
			switch option {
			case "cow":
				appendToAnimalArr(&animalArr, Cow{name})
			case "bird":
				appendToAnimalArr(&animalArr, Bird{name})
			case "snake":
				appendToAnimalArr(&animalArr, Snake{name})
			}
			// animal has to be there in the array to test this.
		case "query":
			var animal Animal
			var err error
			animal, err = findAnimalByName(animalArr, name)
			if err != nil {
				fmt.Println(err)
				break
			}
			switch option {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		}
	}
}
