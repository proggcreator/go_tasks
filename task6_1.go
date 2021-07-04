package main

import (
"fmt"
"sync"
)

func main() {
	var wg sync.WaitGroup
	chFlag := make(chan bool)
	wg.Add(1) //группв из одной горутины
	go func() {
		defer wg.Done()
		for b := range chFlag {
			fmt.Printf("Hello %t\n", b)
		}
	}()
	chFlag <- true
	chFlag <- true
	close(chFlag)
	wg.Wait()
}