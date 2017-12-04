package main

import "fmt"

func main() {
	var queue = make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue)

	//To iterate elements in channel use range.
	for elem := range queue{
		fmt.Println(elem)
	}

}
