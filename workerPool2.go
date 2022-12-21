package main

/*
type ringBuffer struct {
	inCh  chan int
	outCh chan int
}

func NewRingBuffer(inCh, outCh chan int) *ringBuffer {
	return &ringBuffer{
		inCh:  inCh,
		outCh: outCh,
	}
}

func (r *ringBuffer) Run() {
	for v := range r.inCh {
		select {
		case r.outCh <- v:
		default:
			<-r.outCh
			r.outCh <- v

		}
	}
	close(r.outCh)
}

func main() {
	inCh := make(chan int)
	outCh := make(chan int, 10)

	rb := NewRingBuffer(inCh, outCh)

	go rb.Run()

	for i := 0; i < 10; i++ {
		inCh <- i
	}

	close(inCh)

	for r := range outCh {
		fmt.Println(r)
	}
}
*/
