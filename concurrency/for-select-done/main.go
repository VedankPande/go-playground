package main

import (
	"fmt"
	"time"
)


/* 
 For select done - useful when you have an infinitely running task that you want to be able to remotely exit
*/

func worker(done <-chan bool) {
	//infinite loop
	for {
		//select runs task and quits if anything is sent on the done channel
		select {
		case <-done:
			{
				fmt.Println("stopped")
				return
			}
		default:
			{
				fmt.Println("doing work")
			}
		}
	}
}

func main() {
	done := make(chan bool)

	// run worker in goroutine so it doesn't block the main goroutine
	go worker(done)

	//stop worker after 2 seconds
	time.Sleep(time.Second * 2)
	done <- true

}
