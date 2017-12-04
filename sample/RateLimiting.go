package main

import (
	"time"
	"fmt"
)

func main() {

	//Make requests.
	requests := make(chan int, 5)
	for i := 1; i <=5 ; i++ {
		requests <- i
	}
	close(requests)

	//Make rate limiter. Intented to serve one request for every 200 milliseconds.
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests{
		<-limiter
		fmt.Println(req, time.Now())
	}



}
