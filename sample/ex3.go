package main

import (
	"fmt"
	"strconv"
)

func sample(s string) {
	var i int = 0
	for i=0;i<10;i++ {
		fmt.Println(s + "\t" + strconv.Itoa(i))
	}


}

func main() {
	go sample("Hello")

	go sample("Hi")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

