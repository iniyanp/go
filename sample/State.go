package main

import (
	"math/rand"
	"sync/atomic"
	"time"
	"fmt"
)

/**
If we want to operate (read/write) on some critical section say (map) follow these ways.

1. create one go routine that owns critical section. (Go1)
2. create one go routine for read operation.
3. create another go routine for write operation.
4. These read/write go routines contact critical section via channels.
5. Go1 should be responsible to handle read/write operations. It should know what to do when read operation comes and when write
   operation comes.
   

 */

//Construct the data format. Go routines talk via channels with this ds.

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	value int
	resp chan bool
}

func main() {

	var readOps uint64 = 0;
	var writeOps uint64 = 0;

	//So our channel is holding a complex ds not simple ds like int.
	reads := make(chan *readOp);
	writes := make(chan  *writeOp);


	//This routine owns the data structure. and it knows what to do when read channel / write channel comes up.
	go func() {
		var state = make(map[int]int);
		for{
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
				fmt.Println(state)
			}
		}


	}();

	//read operation.

	for i := 0; i < 5 ; i++  {

		go func(){

			for{
				read := &readOp{
					key: rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<- read.resp
				atomic.AddUint64(&readOps,1);
				time.Sleep(1 * time.Millisecond);
			}

		}();

	}

	//write opeartion
	for i := 0; i < 5; i++  {
		go func() {

			for{
				write := &writeOp{
					key : rand.Intn(5),
					value :rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp

				atomic.AddUint64(&writeOps,1);
				time.Sleep(1 * time.Millisecond);

			}
		}()
	}

	time.Sleep(1 * time.Second);
	newWrite := atomic.LoadUint64(&writeOps)
	newRead := atomic.LoadUint64(&readOps)

	fmt.Println(newRead)
	fmt.Println(newWrite)



}


