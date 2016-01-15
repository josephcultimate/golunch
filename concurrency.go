package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string) {

	for j := range jobs {
		sleep := time.Duration(rand.Intn(5)+1) * time.Second
		time.Sleep(sleep)

		results <- fmt.Sprintf("worker %v processing job %v slept for %v", id, j, sleep)
	}
}
func main() {

	jobs := make(chan int, 100)
	results := make(chan string, 100)

	for w := 0; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j <= 8; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 0; a <= 8; a++ {
		fmt.Println("Result: ", <-results)
	}
}
