//Dining Philosopher's Problem

package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS *ChopS
	rightCS *ChopS
}

type Permission struct{
	philosopherId int
	message string
}

func main(){
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
	   CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
	   philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}

	var messageBus [5]chan Permission
	for i := 0; i < 5; i++ {
	   messageBus[i] = make(chan Permission)
	}

	go host(messageBus)

	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 5; i++ {
	   go philos[i].eat(&wg, i, messageBus[i])
	}
	wg.Wait()
}

func (p Philo) eat(wg *sync.WaitGroup, i int, messageBus chan Permission) {
	fmt.Printf("Starting Thread %d\n", i)
	var k, reply Permission
	for j:=0; j<3; j++ {
   		k.philosopherId = i
   		k.message = "PERMISSIONTOEAT"
   		messageBus <- k

   		reply = <- messageBus
		if reply.message == "PERMISSIONGRANTED" && reply.philosopherId == i{
			p.rightCS.Lock()
			p.leftCS.Lock()
			fmt.Printf("Starting to eat %d\n", i)
			fmt.Printf("Finishing eating %d\n", i)
			p.leftCS.Unlock()
			p.rightCS.Unlock()
	   		k.philosopherId = i
	   		k.message = "DONEEATING"
			messageBus <- k
   		}
   		fmt.Printf("Philosopher %d has finished eating %d times\n", i, j+1)
   }
   wg.Done()
}

func host(messageBus [5]chan Permission){
	currentlyEating := 0
	totalEatCount := 0
	var k Permission

	for totalEatCount <= 15 {
		select {
			case message, _ := <- messageBus[0]:
				fmt.Println(message)
				fmt.Printf("Currently Eating: %d and Total Eat Count %d\n", currentlyEating, totalEatCount)
				if message.message=="PERMISSIONTOEAT" {
					if currentlyEating <= 2 {
						k.philosopherId = 0
						k.message = "PERMISSIONGRANTED"
						messageBus[0] <- k
						currentlyEating++
					}
				} else if message.message == "DONEEATING" {
					currentlyEating--
					totalEatCount++
				}
			case message, _ := <- messageBus[1]:
				fmt.Println(message)
				fmt.Printf("Currently Eating: %d and Total Eat Count %d\n", currentlyEating, totalEatCount)
				if message.message=="PERMISSIONTOEAT" {
					if currentlyEating <= 2 {
						k.philosopherId = 1
						k.message = "PERMISSIONGRANTED"
						messageBus[1] <- k
						currentlyEating++
					}
				} else if message.message == "DONEEATING" {
					currentlyEating--
					totalEatCount++
				}
			case message, _ := <- messageBus[2]:
				fmt.Println(message)
				fmt.Printf("Currently Eating: %d and Total Eat Count %d\n", currentlyEating, totalEatCount)
				if message.message=="PERMISSIONTOEAT" {
					if currentlyEating <= 2 {
						k.philosopherId = 2
						k.message = "PERMISSIONGRANTED"
						messageBus[2] <- k
						currentlyEating++
					}
				} else if message.message == "DONEEATING" {
					currentlyEating--
					totalEatCount++
				}
			case message, _ := <- messageBus[3]:
				fmt.Println(message)
				fmt.Printf("Currently Eating: %d and Total Eat Count %d\n", currentlyEating, totalEatCount)
				if message.message=="PERMISSIONTOEAT" {
					if currentlyEating <= 2 {
						k.philosopherId = 3
						k.message = "PERMISSIONGRANTED"
						messageBus[3] <- k
						currentlyEating++
					}
				} else if message.message == "DONEEATING" {
					currentlyEating--
					totalEatCount++
				}
			case message, _ := <- messageBus[4]:
				fmt.Println(message)
				fmt.Printf("Currently Eating: %d and Total Eat Count %d\n", currentlyEating, totalEatCount)
				if message.message=="PERMISSIONTOEAT" {
					if currentlyEating <= 2 {
						k.philosopherId = 4
						k.message = "PERMISSIONGRANTED"
						messageBus[4] <- k
						currentlyEating++
					}
				} else if message.message == "DONEEATING" {
					currentlyEating--
					totalEatCount++
				}
		}
	}
}
