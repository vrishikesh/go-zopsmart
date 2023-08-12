package pkg

import (
	"fmt"
	"sync"
)

// fix the below code
var cache = map[int]int{}

func Q3() {
	t := 100
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for i := 0; i < t; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(sq(i, mu))
		}(i)
	}

	wg.Wait()
}

func sq(n int, mu *sync.Mutex) int {
	defer mu.Unlock()
	mu.Lock()
	if v, ok := cache[n]; ok {
		return v
	}
	cache[n] = n * n
	return cache[n]
}
