package main 

import (
	"fmt"
	"sync"
)

/*
	SimpleCounter is a counter without a mutex that will likely lead to race conditions
	Counter is a counter that uses mutex and thus implements pessimistic locking on the count variables to ensure consistency
*/

//Simple counter - without mutex
type SimpleCounter struct{
	name string
	count int
}

func (c *SimpleCounter) increment(){
	c.count++
}

func simpleCounterWorker(counter *SimpleCounter, numberOfIncrements int, done chan <- bool){

	for i:= 0; i<numberOfIncrements; i++{
		counter.increment()
	}

	//channel synchronization
	done <- true
}

func (c *SimpleCounter) String() string{
	return fmt.Sprintf("Counter %v has count %v",c.name,c.count)
}

//Counter with mutex
type Counter struct{
	name string
	mutex sync.Mutex
	count int
}

func (c *Counter) increment(){

	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count += 1
}


func incrementWorker(counter *Counter, numberOfIncrements int, done chan<- bool){
	for i := 0; i<numberOfIncrements; i++{
		counter.increment()
	}

	done <- true
}

func (c *Counter) String() string{
	return fmt.Sprintf("Counter %v has a count of %v",c.name,c.count)
}

func simulateMutexCounter(numWorkers int){

	done := make(chan bool)
	counter := Counter{name: "A", count: 0}

	for i := 0; i < numWorkers; i++{
		//each worker increments 1000 times
		go incrementWorker(&counter,1000,done)
	}

	//ensure each goroutine finishes before exit
	for j := 0; j<numWorkers; j++{
		<-done
	}

	fmt.Println(&counter)
}

func simulateSimpleCounter(numWorkers int){
	done := make(chan bool)
	counter := SimpleCounter{name: "A",count:0}

	for i := 0; i < numWorkers; i++{

		//each worker increments 1000 times
		go simpleCounterWorker(&counter,1000,done)
	}

	//ensure each goroutine finishes before exit
	for j := 0; j<numWorkers; j++{
		<-done
	}

	fmt.Println(&counter)
}

func main(){

		simulateMutexCounter(3)

		//returns inconsistent results because of race conditions
		simulateSimpleCounter(3)
}

