package main

import (
    "fmt"
)

func bubbleSort(input []int) {
    n := len(input)
    swapped := true
    for swapped {
        swapped = false
        for i := 1; i < n; i++ {
            if input[i-1] > input[i] {
                fmt.Println("Swapping")      
                swap(input, i)
                swapped = true
            }
        }
    }
    fmt.Println(input)
}

func swap(input []int, i int){
    input[i], input[i-1] = input[i-1], input[i]
}


func main() {
    fmt.Println("Please enter the number of integers: ")
    var count int
    _, err := fmt.Scanf("%d", &count)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("Please enter %d integers one by one\n", count)

    toBeSorted := make([]int , 0)

    for i := 0; i < count ; i++ {
        var temp int
        _, err := fmt.Scanf("%d", &temp)

        if err != nil {
            fmt.Println(err)
            break
        }

        toBeSorted = append(toBeSorted, temp)
    }

    bubbleSort(toBeSorted)
}