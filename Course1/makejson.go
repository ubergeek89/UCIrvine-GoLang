package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"encoding/json"
)

func main(){
	fmt.Println("Enter your name : ")
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	name := strings.TrimSuffix(str, "\n")

	fmt.Println("Enter your Address : ")
	in = bufio.NewReader(os.Stdin)
	str, _ = in.ReadString('\n')
	address := strings.TrimSuffix(str, "\n")

	myMap := make(map[string]string)
	myMap["Name"] = name
	myMap["Address"] = address

	jsonString, _ := json.Marshal(myMap)

	fmt.Println(string(jsonString))

}