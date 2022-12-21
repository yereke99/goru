package main

/*
func workerEfficent(id int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)
		go func(job int) {
			defer wg.Done()
			fmt.Println("worker", id, "started job", job)
			time.Sleep(time.Second)
			fmt.Println("worker", id, "fnished job", job)

			results <- job * 2
		}(j)
	}
	wg.Wait()
}

func main() {
	const numjobs = 8

	jobs := make(chan int, numjobs)
	results := make(chan int, numjobs)

	for w := 1; w <= 3; w++ {
		go workerEfficent(w, jobs, results)
	}
	for i := 0; i < numjobs; i++ {
		jobs <- i
	}
	close(jobs)
	for j := 0; j < numjobs; j++ {
		<-results
	}
	close(results)
}
*/
