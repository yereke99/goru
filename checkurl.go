package main

/*
func check(url string, ch chan bool) {
	_, err := http.Get(url)
	if err != nil {
		ch <- false
		return
	}
	ch <- true
}

func main() {
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan bool)

	for _, u := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			mu.Lock()
			check(url, ch)
			mu.Unlock()
		}(u)
	}

	for _, u := range urls {
		fmt.Println("The url: ", u, "is ", <-ch)
	}
	wg.Wait()
}
*/
/*
func check(url string, ch chan bool) {
	_, err := http.Get(url)
	if err != nil {
		ch <- false
		return
	}

	ch <- true
}

func main() {
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	ch := make(chan bool)
	for _, u := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			mu.Lock()
			check(url, ch)
			mu.Unlock()
		}(u)
	}

	for _, u := range urls {
		fmt.Println("The url: ", u, "is ", <-ch)
	}
	wg.Wait()
}
*/
