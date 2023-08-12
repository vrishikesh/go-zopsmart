package pkg

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

func CheckPrimeNumbers() {
	defer TimeTrack(time.Now())
	task := make(chan int)
	out := make(chan string)
	workers := int32(4)

	for i := int32(0); i < workers; i++ {
		go func() {
			defer func() {
				if atomic.AddInt32(&workers, -1) <= 0 {
					close(out)
				}
			}()
			checkPrime(task, out)
		}()
	}

	go func() {
		defer close(task)
		for i := 1; i <= 100000; i++ {
			task <- i
		}
	}()

	for o := range out {
		log.Println(o)
	}
}

func checkPrime(task <-chan int, out chan<- string) {
	for n := range task {
		isPrime := true
		for i := 2; i < n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			out <- fmt.Sprintf("%d is a prime", n)
		}
	}
}
