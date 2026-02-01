package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "processing job", job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// collect results
	for i := 1; i <= 5; i++ {
		fmt.Println("result:", <-results)
	}
}
