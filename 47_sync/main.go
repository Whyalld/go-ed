package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		counter int
		wg sync.WaitGroup
		mutex sync.Mutex
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		
		go func() {
			defer wg.Done()
			increment(&counter, &mutex)
		}()
	}

	wg.Wait()
	fmt.Println("Результат:", counter)
}

func increment(counter *int, mutex *sync.Mutex) {
	mutex.Lock()
	*counter++
	mutex.Unlock()
}