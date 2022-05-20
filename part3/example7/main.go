package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, s := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			fmt.Println(s)
		}(s)
	}
	wg.Wait()
}
