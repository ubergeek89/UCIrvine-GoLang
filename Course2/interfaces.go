package main

import(
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food string
	locomotion string
	noise string
}

type Bird struct {
	food string
	locomotion string
	noise string
}

type Snake struct {
	food string
	locomotion string
	noise string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func main(){
	var command string
	var data map[string]Animal
	data = make(map[string]Animal)
	in := bufio.NewReader(os.Stdin)

	for{
		fmt.Printf("> ")
		str, _ := in.ReadString('\n')
		command = strings.TrimSuffix(str, "\n")
		s := strings.Split(command, " ")
		if(s[0] == "newanimal"){
			animalName := s[1]
			animalType := s[2]
			switch animalType {
				case "cow":
					var animalStruct Cow
					animalStruct.food = "grass"
					animalStruct.locomotion = "walk"
					animalStruct.noise = "moo"
					data[animalName] = animalStruct
					fmt.Println("Created It !")
				case "bird":
					var animalStruct Bird
					animalStruct.food = "worms"
					animalStruct.locomotion = "fly"
					animalStruct.noise = "peep"
					data[animalName] = animalStruct
					fmt.Println("Created It !")
				case "snake":
					var animalStruct Cow
					animalStruct.food = "mice"
					animalStruct.locomotion = "slither"
					animalStruct.noise = "hiss"
					data[animalName] = animalStruct
					fmt.Println("Created It !")
				default:
					fmt.Println("Invalid Animal Type")
			}
		} else if (s[0] == "query") {
			animalName := s[1]
			animalInfo := s[2]
			animalObject , err := data[animalName]
			if err == false {
				fmt.Println("This animal does not exist")
			} else {
				switch animalInfo{
					case "eat":
						animalObject.Eat()
					case "move":
						animalObject.Move()
					case "speak":
						animalObject.Speak()
				}
			}
		} else {
			fmt.Println("Invalid Command")
		}

	}
}