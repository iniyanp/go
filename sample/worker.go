package main

import (
	"fmt"
	"time"
)

//worker id
func myFunc(id int, jobs <-chan int, results chan<- int){
	//range over job channel.
	for p:= range jobs{

		fmt.Println("worker id ", id, "started job ", p)
		time.Sleep(1 * time.Second)
		fmt.Println("worker id ", id, "stopped job ", p)
		//Write the results into result channel.
		results <- p * 2

	}
}

func main() {

	jobs := make(chan int, 10)
	results := make(chan int, 5)

	//call myFunc 5 times using go routine.
	for i:=0;i<5;i++{
		//Only for anonymous function function we need to have ()
		go myFunc(i,jobs,results)
	}

	for i:=0;i<5;i++{
		jobs <- i
	}

	close(jobs) //close the channel indicating that thats all the work we have.

	//Wait for the results. otherwise it will end without printing anything.

	for i:=0;i < 5;i++{
		p := <-results
		fmt.Println("result is ", p)
	}

}
