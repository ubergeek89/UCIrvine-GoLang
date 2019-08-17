package main

import(
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	fn := func (t float64) float64 {
		return 0.5 * a * math.Pow(t,2) + v0 * 2 + s0
	}
	return fn
}

func main(){
	var a, v0, s0, t float64
	fmt.Printf("Enter Acceleration: ")
	fmt.Scanf("%f", &a)
	fmt.Printf("Enter Initial Velocity: ")
	fmt.Scanf("%f", &v0)
	fmt.Printf("Enter Displacement: ")
	fmt.Scanf("%f", &s0)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Printf("Enter Time: ")
	fmt.Scanf("%f", &t)

	fmt.Println(fn(t))
}