package main

/*
import (
	"fmt"
	"sync"
)

func Print(T interface{}) {
	fmt.Println(T)
}
func sum(a []int, c chan int, wg *sync.WaitGroup) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	a := []int{17, 12, 18, 9, 24, 42, 64, 12, 68, 82, 57, 32, 9, 2, 12, 82, 52, 34, 92, 36}
	c := make(chan int)

	for i := 0; i < len(a); i = i + 5 {
		wg.Add(1)
		go sum(a[i:i+5], c, &wg)
	}

	out := make([]int, 5)
	for i := 0; i < 4; i++ {
		out[i] = <-c
	}
	wg.Wait()
	close(c)
	Print(out)

}
*/
/*
func sum(a []int, c chan int, wg *sync.WaitGroup) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	a := []int{17, 12, 18, 9, 24, 42, 64, 12, 68, 82, 57, 32, 9, 2, 12, 82, 52, 34, 92, 36}
	c := make(chan int)

	for i := 0; i < len(a); i = i + 5 {
		wg.Add(1)
		go sum(a[i:i+5], c, &wg)
	}

	out := make([]int, 5)
	for i := 0; i < 4; i++ {
		out[i] = <-c
	}

	wg.Wait()
	close(c)
	fmt.Println(out)
}
*/
