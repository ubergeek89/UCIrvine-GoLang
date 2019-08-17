package main

import "fmt"

func main() {
	fmt.Printf("Enter a floating point number: ")
	var f float64
	_, err := fmt.Scanf("%f", &f)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Truncated Version: %d\n", int(f))
}