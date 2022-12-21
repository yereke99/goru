package main

/*
func sumofSquares(c, quit chan int) {
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return

		}
	}
}

func main() {
	myChanel := make(chan int)
	quit := make(chan int)

	sum := 0

	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-myChanel
		}
		fmt.Println(sum)
		quit <- 0
	}()

	sumofSquares(myChanel, quit)
}
*/
