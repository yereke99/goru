package main

/*
func main() {
	d := time.NewTicker(2 * time.Second)

	myChanel := make(chan bool)

	go func() {
		time.Sleep(30 * time.Second)

		myChanel <- true
	}()

	for {
		select {
		case <-myChanel:
			fmt.Println("Completed!")
			return
		case tm := <-d.C:
			fmt.Println("The current time: ", tm)
		}
	}
}
*/
