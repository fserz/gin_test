package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 1; i <= 5000; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() {
		for i := 5000; i <= 100000; i++ {
			c2 <- i
		}
		close(c2)
	}()

	merged := merge(c1, c2)

	for val := range merged {
		fmt.Println(val)
	}
	fmt.Println("Done")
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	// 不开新的 goroutine 会死锁
	//wg.Wait()
	//close(out)
	return out
}
