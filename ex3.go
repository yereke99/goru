package main

/*
// Working with chanels
func squares(c chan int) {
	for i := 1; 1 <= 5; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)

	fmt.Println("active goruines", runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	fmt.Println("active goruines", runtime.NumGoroutine())

	go squares(c)
	fmt.Println("active goruines", runtime.NumGoroutine())

	c <- 5
	c <- 6
	c <- 7
	c <- 8

	fmt.Println("active goroutines", runtime.NumGoroutine())
	fmt.Println("main() stopped")

}
*/
/*
func squares(c chan int) {
	for i := 1; i <= 5; i++ {
		num := <-c
		fmt.Println(num * num)
	}

}

func main() {
	fmt.Println("main() started")

	c := make(chan int, 2)

	go squares(c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5

	fmt.Println("main() stoped")

	//time.Sleep(time.Second * 2)
}
*/
/*
func squares(c chan int) {
	for i := 1; i <= 9; i++ {
		c <- i * i
	}

	close(c)
}

func main() {
	fmt.Println("main() started")
	c := make(chan int)

	go squares(c)

	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "<-- loop broke!")
			break
		} else {
			fmt.Println(val, ok)
		}
	}
	fmt.Println("main() stoped")
}

*/
/*
func squares(c chan int) {
	for i := 1; i <= 9; i++ {
		c <- i * i
	}

	close(c)
}

func main() {
	fmt.Println("main() started")
	c := make(chan int)

	go squares(c)

	for val := range c {
		fmt.Println(val)
	}
	fmt.Println("main() stoped")

}
*/
