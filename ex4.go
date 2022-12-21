package main

/*
// Mutex бұғаттауды қолдана отыра санды көбейту
var n = 1

var mu sync.Mutex

func timesThree() {
	mu.Lock()
	defer mu.Unlock()
	n *= 3
	fmt.Println(n)
}

func main() {
	fmt.Println("We are executing a goroutine")

	for i := 0; i <= 10; i++ {
		go timesThree()
	}

	time.Sleep(time.Second)
}
*/
/*
// Санды 3ке көбейту.
var n = 1

func timesThree() {
	n *= 3
	fmt.Println(n)
}

func main() {
	fmt.Println("We are executing a goroutine")

	for i := 0; i <= 10; i++ {
		go timesThree()
	}

	time.Sleep(time.Second)
}
*/

/*
// Екі каналдымен жұмыс жасау
func timeThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem * 3
	}
}

func minusThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem - 3
	}
}

func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{2, 3, 4, 5, 6}
	ch := make(chan int, len(arr))
	minusCh := make(chan int, len(arr))

	go timeThree(arr, ch)
	go minusThree(arr, minusCh)

	for i := 0; i < len(arr)*2; i++ {
		select {
		case msg1 := <-ch:
			fmt.Printf("Result timesThree: %v \n", msg1)
		case msg2 := <-minusCh:
			fmt.Printf("Result minusThree: %v \n", msg2)

		default:
			fmt.Println("Non blocking way of listening to multiple channels")
		}
	}
}
*/

/*
func timesThree(arr []int, ch chan int) {
	minusCh := make(chan int, 9)
	for _, elem := range arr {
		value := elem * 3

		if value%2 == 0 {
			go minusThree(value, minusCh)
			value = <-minusCh
		}

		ch <- value

	}
}

func minusThree(number int, ch chan int) {
	ch <- number - 3
	fmt.Println("The functions continues after returning the result")
}

func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan int, len(arr))

	go timesThree(arr, ch)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Result: %v\n", <-ch)
	}
}
*/
/*
func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	ch := make(chan int, len(arr))

	go func(arr []int, ch chan int) {
		for _, elem := range arr {
			ch <- elem * 3
		}
	}(arr, ch)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Result: %v \n", <-ch)
	}

}
*/
