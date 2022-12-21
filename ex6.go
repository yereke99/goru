package main

/*
import (
	"fmt"
	"sync"
)

var counter int = 0

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup

	result := make(chan int)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(ch chan int, m *sync.Mutex, wg *sync.WaitGroup) {
			m.Lock()
			counter += 1
			m.Unlock()
			ch <- counter
			wg.Done()
		}(result, &m, &wg)
	}

	for i := 0; i < 1000; i++ {
		fmt.Println("Result: ", <-result)
	}
	wg.Wait()

}
*/
