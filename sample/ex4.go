package main

import (
	"time"
	"fmt"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Wait for the timer1 is over")


}
