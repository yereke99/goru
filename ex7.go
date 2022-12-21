package main

/*
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// naturals
	go func() {
		for x := 0; x < 1000; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// squares
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
*/
/*
import "fmt"

func main() {
	natural := make(chan int)
	squares := make(chan int)

	// naturals
	go func() {
		for x := 0; x < 1000; x++ {
			natural <- x
		}
		close(natural)
	}(Println()

	// squares
	go func() {
		for x := range natural {

			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
*/
