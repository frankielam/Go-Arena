package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", job)
		results <- job
	}

}

func main() {
	const jobNum = 5 
	jobs := make(chan int, jobNum)
	results := make(chan int, jobNum)

	for i := 0; i < jobNum; i++ {
		go worker(i, jobs, results)	
	}

	for j := 0; j < jobNum; j++ {
		jobs <- j
	}
	close(jobs)


	for k := 0; k < jobNum; k++ {
		fmt.Println( <-results)
	}

}