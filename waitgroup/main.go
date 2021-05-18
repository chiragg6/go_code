package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("Running go routine")
		defer wg.Done()
	}()

	wg.Wait()
}
