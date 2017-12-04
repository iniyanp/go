package main

import (
	"fmt"
	"time"
)

func normalChannel() {
	//create a channel.
	messages := make(chan string)
	//go routine.
	go func() {messages <- "Iniyan is a good boy"}()

	//Automatically wait for the message to receive. no need to put extra synchronization.
	var p1 = <- messages
	fmt.Println(p1)
}

func bufferedChannel() {
	//Buffered channel.

	//size is 2. so without corresponding receive immediately, it can wait.
	ch1 := make(chan string, 2)

	ch1 <- "Kumar"
	ch1 <- "Thamarai"

	fmt.Println(<- ch1);
	fmt.Println(<- ch1);
}

func worker(c chan bool) {
	fmt.Println("Starting")
	time.Sleep(3 * time.Second)
	fmt.Println("done")
	c <- true
}

func syncChannel() {

	var c = make(chan bool, 1)
	go worker(c) //call worker method using go routine.
	<-c //This is very important. so that you are waiting here .
}

//while receiving the channel parameter in function, we can mention whether this channel is meant for sending or receiving.


//chan
//<-chan get the value from channel.
//chan<- send value to channel.

//This specificity increases the type-safety of the program.
//Combining goroutines and channels with select is a powerful feature of Go.

func ping(pings chan<- string, msg string){
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string){
	//Whenever you need to extract something from the channel, use <-
	msg := <-pings
	pongs <- msg
}

func pingPong() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings,"Hello")
	pong(pings,pongs)
	fmt.Println(<-pongs)
}

func main() {

	//var p = []int{1,2,3} //slice
	//
	////range is used to iterate slice, array, struct.. etc.
	//for i := range p {
	//	fmt.Println(p[i])
	//}
	//
	//var m = map[string] string{"Iniyan" : "1","Kumar":"2"}
	//for key := range m{
	//	fmt.Println(key, m[key])
	//}

	//bufferedChannel()
	//syncChannel()
	pingPong()




}
