package main

/*
import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {
	time.Sleep(250 * time.Microsecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"

	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}

		}
	}

}

func main() {
	var wg sync.WaitGroup
	cmd := make(chan string)
	wg.Add(1)

	go routine(cmd, &wg)

	time.Sleep(1 * time.Second)
	cmd <- "Pause"

	time.Sleep(1 * time.Second)
	cmd <- "Play"

	time.Sleep(1 * time.Second)
	cmd <- "Stop"

	wg.Wait()

	close(cmd)
}
*/
