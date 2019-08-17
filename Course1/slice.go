package main

import(
	"fmt"
	"strconv"
	"sort"
)

func main(){
	sli := make([]int , 0)
	var x string;
	for{
		fmt.Scan(&x)
		if x == "X" || x == "x" {
			break
		}
		y, err := strconv.Atoi(x)

		if err != nil {
			fmt.Println("error")
			break
		}
		sli = append(sli, y)
		sort.Ints(sli)
		fmt.Println(sli)
	}
}