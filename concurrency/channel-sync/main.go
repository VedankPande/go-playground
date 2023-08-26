package main

import (
	"fmt"
	"time"
)

/*
channel synchronization to wait for a goroutine to finish before the main function is exited
*/
func main() {

	done := make(chan bool)

	//anonymous function that takes a receive directed channel as input
	go func(done chan<- bool) {

		fmt.Println("Started task")
		time.Sleep(time.Second * 2)
		fmt.Println("Task complete!")

		done <- true
	}(done)
	
	//the done channel here is blocked until it receives a value (in this case from the goroutine above)
	<-done
}
