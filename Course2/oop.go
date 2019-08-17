package main

import(
	"fmt"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat(){
	fmt.Println(a.food)
}

func (a Animal) Move(){
	fmt.Println(a.locomotion)
}

func (a Animal) Speak(){
	fmt.Println(a.noise)
}
func main(){
	var cow Animal
	cow.food = "grass"
	cow.locomotion = "walk"
	cow.noise="moo"

	var bird Animal
	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noise="peep"

	var snake Animal
	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noise="hsss"

	for{
		var animal, info string
		fmt.Printf("> ")
		fmt.Scan(&animal)
		fmt.Printf("> ")
		fmt.Scan(&info)
		switch {
			case animal == "cow" && info == "eat":
				cow.Eat()
			case animal == "cow" && info == "move":
				cow.Move()
			case animal == "cow" && info == "speak":
				cow.Speak()
			case animal == "bird" && info == "eat":
				bird.Eat()
			case animal == "bird" && info == "move":
				bird.Move()
			case animal == "bird" && info == "speak":
				bird.Speak()
			case animal == "snake" && info == "eat":
				snake.Eat()
			case animal == "snake" && info == "move":
				snake.Move()
			case animal == "snake" && info == "speak":
				snake.Speak()
			default:
				fmt.Println("Not recognized")
		}
	}
}