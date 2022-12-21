package main

/*

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.v[key]++
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("Go")
		go c.Inc("Python")
	}

	for i := 0; i <= 1000; i++ {
		go c.Inc("JavaScript")
		go c.Inc("Python")
		go c.Inc("Python")
	}

	time.Sleep(2 * time.Second)
	fmt.Println(c.Value("Go"))
	fmt.Println(c.Value("Python"))
	fmt.Println(c.Value("JavaScript"))
}
*/

/*
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	s.v[key]++
	s.mu.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go c.Inc("SomeKey")
	}

	for i := 0; i <= 1000; i++ {
		go c.Inc("SomeKeys")
	}

	time.Sleep(2 * time.Second)
	fmt.Println(c.Value("SomeKey"))
	fmt.Println(c.Value("SomeKeys"))

}
*/
