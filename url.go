package main

/*
func checkUrls(urls []string) {
	c := make(chan string)

	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)

		go func(wg *sync.WaitGroup, url string, c chan string) {
			defer wg.Done()

			_, err := http.Get(url)

			if err == nil {
				c <- "Success reaching the website: " + url
			} else {
				c <- "We could not search: " + url
			}
		}(&wg, u, c)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	start := time.Now()
	links := []string{
		"https://go.dev/learn/",
		"https://www.netflix.com/",
		"https://go.dev/learn/",
		"https://www.netflix.com/",
		"https://go.dev/learn/",
		"https://www.netflix.com/",
		"https://github.com/",
	}

	checkUrls(links)
	fmt.Printf("Completed the code process, took: %f seconds\n", time.Since(start).Seconds())
}
*/
