// This program is non-deterministic and will produce a deffierent output whenever it is run.

package main

import (
	"fmt"
)

//Global variable
var intCounter int

//Main function is also a goroutine
func main() {
	//Sets intcounter from 1 to 50 and starts a thread at each iteration
	for intCounter = 0; intCounter < 50; intCounter++ {
		go routineone()
	}

	//The main routine ends much before all the goroutines finishes.
	//Only some of the goroutines end up executing the print statement.
}

func routineone(){
	fmt.Println(intCounter)
}