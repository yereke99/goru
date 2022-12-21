package main

/*
func worker(id int, wg *sync.WaitGroup, jobs <-chan int, result chan<- int) {
	defer wg.Done()

	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	var wg sync.WaitGroup
	const jobN = 5

	jobs := make(chan int, jobN)
	results := make(chan int, jobN)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, &wg, jobs, results)
	}

	for j := 1; j <= jobN; j++ {
		jobs <- j
	}

	close(jobs)
	for a := 1; a <= jobN; a++ {
		<-results
	}
	wg.Wait()
}
*/
/*
func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	var wg sync.WaitGroup
	const jobN = 5

	jobs := make(chan int, jobN)
	results := make(chan int, jobN)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, &wg, jobs, results)
	}

	for j := 1; j <= jobN; j++ {
		jobs <- j
	}

	close(jobs)
	for a := 1; a <= jobN; a++ {
		<-results
	}
	wg.Wait()

}
*/
/*
func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	var wg sync.WaitGroup
	const NumJobs = 5
	jobs := make(chan int, NumJobs)
	results := make(chan int, NumJobs)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, &wg, jobs, results)
	}

	for j := 1; j <= NumJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= NumJobs; a++ {
		<-results
	}
	wg.Wait()

}
*/
