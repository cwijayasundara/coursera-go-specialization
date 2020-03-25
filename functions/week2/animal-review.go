package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func AnimalMap() map[string]Animal {
	m := make(map[string]Animal)
	m["cow"] = Animal{"grass", "walk", "moo"}
	m["bird"] = Animal{"worms", "fly", "peep"}
	m["snake"] = Animal{"mice", "slither", "hsss"}
	return m
}

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Bad prompt read! %d, %s\n", s, err)
		os.Exit(-1)
	} else {
		s = strings.TrimRight(s, "\n")
	}
	return s
}

func main() {
	animalFarm := AnimalMap()

	for {
		s := prompt()
		var a, b string
		n, err := fmt.Sscanf(s, "%s %s\n", &a, &b)
		if n != 2 || err != nil {
			fmt.Printf("Bad conversion! %d, %s\n", n, err)
			continue
		}

		c, bb := animalFarm[a]
		if !bb {
			fmt.Printf("Bad animal: %s. Valid animals are 'cow', 'bird', 'snake'\n", a)
			continue
		}
		switch b {
		case "eat":
			c.Eat()
		case "move":
			c.Move()
		case "speak":
			c.Speak()
		default:
			fmt.Println("Bad action: %s. Valid actions are 'eat', 'move' or 'speak'\n", b)
		}
	}
}
