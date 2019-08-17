package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

func main(){
    fmt.Println("Please enter the number of integers to sort: ")
    var count int
    _, err := fmt.Scanf("%d", &count)

    if err != nil {
        fmt.Println(err)
    }

    size1 := int(math.Round(float64(count)/float64(4)))
    size2, size3 := size1, size1
    size4 := count - (size1 + size2 + size3)
    //fmt.Println(size1, size2, size3, size4)

    slice1 := make([]int , 0)
    slice2 := make([]int , 0)
    slice3 := make([]int , 0)
    slice4 := make([]int , 0)

    fmt.Printf("Please enter %d integers one by one (followed by return key) : \n", count)

    for i := 0; i < size1 ; i++ {
        var temp int
        _, err := fmt.Scanf("%d", &temp)
        if err != nil {
            fmt.Println(err)
            break
        }
        slice1 = append(slice1, temp)
    }

    for i := 0; i < size2 ; i++ {
        var temp int
        _, err := fmt.Scanf("%d", &temp)
        if err != nil {
            fmt.Println(err)
            break
        }
        slice2 = append(slice2, temp)
    }

    for i := 0; i < size3 ; i++ {
        var temp int
        _, err := fmt.Scanf("%d", &temp)
        if err != nil {
            fmt.Println(err)
            break
        }
        slice3 = append(slice3, temp)
    }

    for i := 0; i < size4 ; i++ {
        var temp int
        _, err := fmt.Scanf("%d", &temp)
        if err != nil {
            fmt.Println(err)
            break
        }
        slice4 = append(slice4, temp)
    }

    fmt.Println(slice1, slice2, slice3, slice4)

    fmt.Println("Starting 4 goroutines")

	var wg sync.WaitGroup
	wg.Add(4)
	go sortMethod(&wg, slice1)
	go sortMethod(&wg, slice2)
	go sortMethod(&wg, slice3)
	go sortMethod(&wg, slice4)
	wg.Wait()

    fmt.Println("Goroutines ended")

    fmt.Println(slice1, slice2, slice3, slice4)

    fmt.Println("Let us now merge all the 4 slices")

    sortedArray1 := merge(slice1, slice2)
    sortedArray2 := merge(slice3, slice4)
    sortedArray := merge(sortedArray1, sortedArray2)


    fmt.Println("\n\nFinal Result\n===========")
    fmt.Println(sortedArray)

}

func sortMethod(wg *sync.WaitGroup, slice []int){
	fmt.Printf("I am sorting the slice %v\n", slice)
	sort.Ints(slice)
	wg.Done()
}

func merge(left, right []int) (result []int) {
    result = make([]int, len(left) + len(right))
      
    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }
      
    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }
      
    return
}