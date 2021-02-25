/*	Copyright 2021 www.0example.com, powered by Wixis360  */


package main

import (
	"fmt"
)
import "time"

//We run several  concurrent instances of the worker.
//Each worker will receive work on jobs channel and send the relevent result on reuslts channel.
//Function Sleep is used to simulate the work as an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker:", id, "Started Job : ", j)
		time.Sleep(time.Second)
		fmt.Println("Worker:", id, "Finished Job : ", j)
		results <- j * 2

		fmt.Println()
	}
}

func main() {

	//Two channels are created
	//job channel - send work to workers
	// work channel - receive the result of the work.
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Two workers are set but initially blocked until jobs are received to the channel
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	//Here we send 4 jobs to the channel
	for j := 1; j <= 4; j++ {
		jobs <- j
	}

	//Then we close the channel after all the work put into it, 
	//to signal that's all the work we have.
	close(jobs)

	//At the end, we collect all the results of the work.
	for a := 1; a <= 4; a++ {
		<-results
	}
}

//After running the programme we can see, even though it takes about 4 seconds to complete the 4 tasks,
//it only took about 2 seconds to finish as the workers did the tasks concurrently