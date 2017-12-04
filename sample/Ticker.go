package main

import (
	"time"
	"fmt"
)

//If we want to execute something at regular intervals, we can use Ticker.
func printme(elem time.Time) {
	fmt.Println("Iniyan " + elem.String())
}

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for elem := range ticker.C {
			printme(elem)
		}
	}()
	time.Sleep(20 * time.Second)
	ticker.Stop()

}
