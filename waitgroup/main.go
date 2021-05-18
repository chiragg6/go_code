package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)
	go func(ch chan string) {
		fmt.Println("Running go routine")
		defer wg.Done()
		ch <- "testing channels"
	}(ch)
	a := <-ch
	fmt.Println("Printing value from channel", a)
	wg.Wait()
}
