package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 10000000; i++ {
		fn()
	}
}

func fn() {
	var memoryAccess sync.Mutex
	var value int
	var s string

	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0 {
		s = fmt.Sprintf("value:%d", value)
	}
	memoryAccess.Unlock()

	// if s == "value:0" {
	// 	fmt.Println("RUN")
	// }
	// if s == "value:1" {
	// 	fmt.Println("RUN")
	// }
	if s == "" {
		fmt.Println("RUN")
	}
}
