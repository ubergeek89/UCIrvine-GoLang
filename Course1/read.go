package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
)

type Person struct{
	fname string
	lname string
}

func main(){
	fmt.Println("Enter file name : ")
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	fileName := strings.TrimSuffix(str, "\n")

	file, err := os.Open(fileName)
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()

	sli := make([]Person , 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	    s := scanner.Text()
	    stringarr := strings.Split(s, " ")
	    fname, lname := stringarr[0], stringarr[1]
	    var obj Person
	    obj.fname = fname
	    obj.lname = lname
	    sli = append(sli, obj)
	}

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}

	for _,val := range sli {
		fmt.Printf("%s %s\n", val.fname, val.lname)
	}

}