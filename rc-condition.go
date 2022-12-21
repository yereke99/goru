package main

/*
var G = 0

func Worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	G += 1

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go Worker(&wg, &m)
	}

	wg.Wait()
	fmt.Println("Value of x: ", G)
}
*/
