package main

import (
	"time"
	"sync/atomic"
	"fmt"
)

func main() {


	var ops uint64 = 0
	for i := 0; i < 2 ; i++  {
		go func() {

			//If we need to increment atmoic counter, we have to use for
			for{
				atomic.AddUint64(&ops,1)
				time.Sleep(500 * time.Millisecond);
			};

		}();

	}

	//You need to wait and then take a copy of that variable.
	time.Sleep(1000 * time.Millisecond);

	opsFinal := atomic.LoadUint64(&ops);
	fmt.Println(opsFinal);
}


