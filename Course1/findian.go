package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func main(){
	var s string
	fmt.Printf("Please enter a string: ")

	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	s = strings.TrimSuffix(str, "\n")

    if err == nil {
      //
	}
	s = strings.ToLower(s)

	if strings.Index(s, "i") == 0 {
		if strings.Index(s, "n") == len(s)-1 {
			if strings.Index(s, "a") > 0 && strings.Index(s,"a") < len(s) -1 {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}
		} else {
			fmt.Println("Not Found!")
		}		
	} else {
		fmt.Println("Not Found!")
	}

}