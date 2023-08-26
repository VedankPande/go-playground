package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
Worker pools limit the number of workers (goroutines) that can be run at a time to conserve resources.
Each worker polls for a task and returns to the worker pool once it is done executing the task
*/
func worker(workerID int, tasks <-chan string, results chan<- string) {
	for task := range tasks {

		fmt.Println("worker:", workerID, "Started task:", task)
		time.Sleep(time.Second * 2)
		results <- fmt.Sprintf("worker %v completed task %v", workerID, task)
	}
}

func main() {
	//number of jobs/tasks and workers in the pool
	numTasks := 10
	numWorkers := 3

	//channels for tasks and returning results
	tasks := make(chan string, numTasks)
	results := make(chan string, numTasks)

	// NOTE: if you spin up workers after creating tasks, you'll notice a change in order of workers starting tasks

	//create workers
	for w := 0; w < numWorkers; w++ {
		go worker(w, tasks, results)
	}
	
	//send tasks to the channel
	for i := 0; i < numTasks; i++ {
		tasks <- fmt.Sprintf("Task %v", strconv.Itoa(i+1))
	}

	close(tasks)

	//get results
	for r := 0; r < numTasks; r++ {
		fmt.Println(<-results)
	}

}
