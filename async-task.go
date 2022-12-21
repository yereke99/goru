package main

/*
import (
	"fmt"
	"sync"
	"time"

	cron "github.com/robfig/cron"
)

type Visited struct {
	mux     sync.Mutex
	rMux    sync.RWMutex
	visited map[string]int
}

func (v *Visited) Inc(url string) {
	v.mux.Lock()
	defer v.mux.Unlock()

	v.visited[url]++
}

func (v *Visited) Value(url string) int {
	v.rMux.RLock()
	defer v.rMux.RUnlock()
	return v.visited[url]
}

func main() {
	s1 := cron.New()
	s2 := cron.New()
	s3 := cron.New()

	s := Visited{visited: make(map[string]int)}

	s1.AddFunc("@every 0h0m1s", func() {
		for {
			time.Sleep(100 * time.Microsecond)
			s.Inc("google.kz")
		}
	})

	s2.AddFunc("@every 0h0m1s", func() {
		for {
			time.Sleep(10 * time.Microsecond)
			s.Inc("google.com")
		}
	})

	s3.AddFunc("@every 0h0m1s", func() {
		for {
			time.Sleep(10 * time.Microsecond)
			s.Inc("amazon.com")
		}
	})

	defer s1.Stop()
	defer s2.Stop()
	defer s3.Stop()

	go s1.Start()
	go s2.Start()
	go s3.Start()
	time.Sleep(10 * time.Second)

	fmt.Println(s.Value("google.kz"))
	fmt.Println(s.Value("google.com"))
	fmt.Println(s.Value("amazon.com"))
}
*/
