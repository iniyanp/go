package main

import (
	"time"
	"fmt"
)

func func1(c1 chan string) {

	time.Sleep(3 * time.Second)
	c1 <- "Hello"
}

func func2(c2 chan string) {

	time.Sleep(3 * time.Second)
	c2 <- "How are you?"
}

func main() {

	var c1 = make(chan string,1)
	var c2 = make(chan string,1)

	go func1(c1)
	go func2(c2)

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout from func1")
	}

	//you are waiting for c2.
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout from func2")
	}



}
