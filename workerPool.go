package main

/*
var (
	mu      sync.Mutex
	counter int
)

func work(n int) <-chan int {
	r := make(chan int)

	for i := 0; i < n; i++ {
		go func() {
			for c := 0; c < 100000; c++ {
				mu.Lock()
				counter += 1
				mu.Unlock()
			}
			r <- counter
			close(r)
		}()
		close(r)
	}

	return r
}

func Print(c <-chan int) {
	fmt.Println(<-c)
}

func main() {
	c := work(5)
	Print(c)
}
*/
/*
var (
	wg      sync.WaitGroup
	counter int
)

func work(n int) <-chan int {
	r := make(chan int)
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 1e10; i++ {
				counter++
			}
			r <- counter
			close(r)
		}()
	}

	return r

}

func Print(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func main() {

	c := work(5)

	Print(c)

}
*/
