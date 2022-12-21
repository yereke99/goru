package main

/*
func timesThree(n int, c chan int) {
	result := n * 3
	fmt.Println(result)
	c <- result
}

func main() {
	fmt.Println("We are executing chanel!")
	ch := make(chan int)
	go timesThree(3, ch)
	result := <-ch
	fmt.Printf("The result is: %v", result)
}

*/

/*
func timesThree(n int, c chan int) {
	result := n * 3
	//fmt.Println(result)
	c <- result
}

func main() {
	fmt.Println("We are executing chanel!")
	ch := make(chan int)
	for i := 1; i <= 10; i++ {
		go timesThree(i, ch)
		result := <-ch
		fmt.Printf("The result is: %v\n", result)
	}

}
*/
/*
func timesThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem * 3
	}
}

func main() {
	fmt.Println("We are executing chanel!")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan int, len(arr))
	go timesThree(arr, ch)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("The result: %v\n", <-ch)
	}

}
*/
/*
func main() {
	fmt.Println("We are executing chanel!")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan int, len(arr))

	go func(arr []int, ch chan int) {
		for _, elem := range arr {
			ch <- elem * 3
		}
	}(arr, ch)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("The result: %v\n", <-ch)
	}
}
*/
/*
func main() {
	go spinner(100 * time.Microsecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
*/

/*
func main() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				ch <- "Gorutine : " + strconv.Itoa(i)
			}
		}(i)
	}

	for k := 1; k <= 10; k++ {
		fmt.Println(k, <-ch)
	}
}
*/
