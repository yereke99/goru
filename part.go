package main

/*
var (
	partList    = []string{"A", "B", "C", "D"}
	nAssemblies = 3
	wg          sync.WaitGroup
)

func worker(part string) {
	defer wg.Done()

	log.Println(part, "worker begins start")
	time.Sleep(time.Microsecond * 1000)
	log.Println(part, "worker completes part")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for c := 1; c <= nAssemblies; c++ {
		log.Println("begin assembly cycle", c)
		wg.Add(len(partList))
		for _, part := range partList {
			go worker(part)
		}

		wg.Wait()
		log.Println("assemble.  cycle", c, "complete")
	}
}
*/
/*
var (
	partList    = []string{"A", "B", "C", "D"}
	nAssemblies = 3
	wg          sync.WaitGroup
)

func worker(part string) {
	log.Println(part, "worker begins start")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	log.Println(part, "worker completes part")
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for c := 1; c <= nAssemblies; c++ {
		log.Println("begin assembly cycle", c)
		wg.Add(len(partList))
		for _, part := range partList {
			go worker(part)
		}
		wg.Wait()
		log.Println("assemble.  cycle", c, "complete")
	}
}
*/
