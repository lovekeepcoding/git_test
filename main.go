package main

import (
	"fmt"
	"sync"
)

var (
	sum int = 0
	lock sync.RWMutex
	wg sync.WaitGroup
)

func Add()  {
	defer wg.Done()
	lock.Lock()
	sum += 10
	lock.Unlock()
}

func main()  {
	wg.Add(100)

	for i := 0; i < 100; i ++ {
		go Add()
	}
	wg.Wait()
	fmt.Println(sum)
}
