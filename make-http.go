package main

/*
func makeHttp(link string) (url chan string, e chan error) {
	url = make(chan string)

	e = make(chan error)

	go func() {
		_, err := http.Get(link)

		if err == nil {
			url <- link
		} else {
			e <- err
		}
		close(url)
		close(e)
	}()

	return
}

func main() {
	links := []string{
		"http://abc.com",
		"http://pqr.com",
		"http://xyz.com",
	}

	for _, link := range links {
		r, err := makeHttp(link)
		fmt.Println(<-r, <-err)
	}
}
*/

/*
func makeHttp(link string) (url chan string, e chan error) {
	url = make(chan string)
	e = make(chan error)
	go func() {
		_, err := http.Get(link)
		if err == nil {
			url <- link
		} else {
			e <- err
		}
		close(url)
		close(e)

	}()

	return
}

func main() {
	links := []string{
		"http://abc.com",
		"http://pqr.com",
		"http://xyz.com",
	}

	for _, link := range links {
		r, err := makeHttp(link)
		fmt.Println(<-r, <-err)
	}

}
*/
