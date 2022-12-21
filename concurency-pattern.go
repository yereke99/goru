package main

/*
func takeFirstN(ctx context.Context, dataSource <-chan interface{}, n int) <-chan interface{} {
	takeChannel := make(chan interface{})
	go func() {
		defer close(takeChannel)

		for i := 0; i < n; i++ {
			select {
			case val, ok := <-dataSource:
				if !ok {
					return
				}
				takeChannel <- val
			case <-ctx.Done():
				return
			}
		}
	}()

	return takeChannel
}

func rangeNum(ctx context.Context, fetchers ...<-chan interface{}) <-chan interface{} {
	combinedFetchers := make(chan interface{})

	var wg sync.WaitGroup

	wg.Wait()

	for _, f := range fetchers {
		f := f
		go func() {
			defer wg.Done()
			for {
				select {
				case res := <-f:
					combinedFetchers <- res
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(combinedFetchers)
	}()

	return combinedFetchers
}

func main() {
	done := make(chan struct{})
	defer close(done)
	//range10 := rangeNum(done, 10)
	//for num := range takeFirstN(done)
}
*/
