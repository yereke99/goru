package main

/*
func Fibo(ch chan int, quite chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quite:
			fmt.Println("Quite")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	q := make(chan bool)
	n := 10
	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch)
		}
	}(n)

	Fibo(ch, q)
}
*/
