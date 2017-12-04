package main

import (
	"time"
	"fmt"
	"strconv"
)

//time go run select.go

func first(c chan string) {
	time.Sleep(2 * time.Second)
	var i int;
	//we can keep sending messages to this. :)
	for i=0;i<2;i++{
		c <- "Iniyan" + strconv.Itoa(i)
	}
}

func second(c chan string) {
	time.Sleep(5 * time.Second)
	var i int;
	for i=0;i<2;i++ {
		c <- "Paramasivam" + strconv.Itoa(i)
	}
}


func main() {
	c1 := make(chan string,1)
	c2 := make(chan string,1)

	go first(c1)
	go second(c2)

	//var p1 = <-c2
	//fmt.Println(p1)

	var i int
	//you are waiting for a channels' message in a loop.
	for i=0;i<4;i++ {
		select {
		//Go’s select lets you wait on multiple channel operations.
		case msg1 := <-c1:
			fmt.Println(msg1)

		case msg2 := <-c2:
			fmt.Println(msg2)

		}
	}

}
