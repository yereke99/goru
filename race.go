package main

/*
var sharedMapForMutex map[string]int = map[string]int{}
var mapMutex = sync.Mutex{}
var mutexReadCount int64 = 0

func runMapMutexReader(ctx context.Context, readChan chan int) {
	readCount := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("reader existing...")
			readChan <- readCount
			return
		default:
			mapMutex.Lock()
			var _ int = sharedMapForMutex["key"]
			mapMutex.Unlock()
			readCount += 1
		}
	}
}
func runMApMutexWriter(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writer existing ...")
			return
		default:
			mapMutex.Lock()
			sharedMapForMutex["key"] = sharedMapForMutex["key"] + 1
			mapMutex.Unlock()
			time.Sleep(100 * time.Microsecond)
		}
	}
}
func startMapMutexReadWrite() {
	testContext, cancel := context.WithCancel(context.Background())

	readChan := make(chan int)
	sharedMapForMutex["key"] = 0
	numberOfReaders := 15
	for i := 0; i < numberOfReaders; i++ {
		go runMapMutexReader(testContext, readChan)
	}

	go runMApMutexWriter(testContext)

	time.Sleep(2 * time.Second)
	cancel()
	totalReadCount := 0

	for i := 0; i < numberOfReaders; i++ {
		totalReadCount += <-readChan
	}

	time.Sleep(1 * time.Second)

	var counter int = sharedMapForMutex["key"]
	fmt.Printf("[MUTEX] Write counter value: %v\n", counter)
	fmt.Printf("[MUTEX] Read counter value: %v\n", totalReadCount)
}

func main() {
	startMapMutexReadWrite()
}
*/
