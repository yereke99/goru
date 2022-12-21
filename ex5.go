package main

/*
func Init() {
	rand.Seed(time.Now().Unix())
}

func Sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func reader(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
	Sleep()
	m.RLock()
	c <- 1
	Sleep()
	c <- -1
	m.RUnlock()
	wg.Done()
}

func writer(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
	Sleep()
	m.Lock()
	c <- 1
	Sleep()
	c <- -1
	m.Unlock()
	wg.Done()
}

func main() {
	var m sync.RWMutex
	var rs, ws int

	rsCh := make(chan int)
	rwCh := make(chan int)

	go func() {
		for {
			select {
			case n := <-rsCh:
				rs += n
			case n := <-rwCh:
				ws += n
			}
			fmt.Printf("%s%s\n", strings.Repeat("R", rs), strings.Repeat("W", ws))
		}
	}()

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go reader(rsCh, &m, &wg)
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go writer(rsCh, &m, &wg)
	}

	wg.Wait()

}
*/
