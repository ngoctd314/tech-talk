package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func execIn1s() {
	now := time.Now()
	for j := 0; j < 4_000_000_000; j++ {
	}
	log.Println("time execute a heavy load function: ", time.Since(now))
}

func main() {
	fmt.Println("===================")
	fmt.Println("GO VER: ", runtime.Version())
	numCpu := os.Getenv("GOMAXPROCS")
	procs, _ := strconv.Atoi(numCpu)
	n := runtime.GOMAXPROCS(procs)
	fmt.Println("num pros:", n)

	execIn1s()
	execIn1s()
	execIn1s()

	// // sequentialVer()
	// concurrentVer(n) // expect time: t*[10/n]
}

func sequentialVer() {
	log.Println("Sequential version")
	now := time.Now()

	for i := 0; i < 20; i++ {
		execIn1s()
	}

	log.Println("since: ", time.Since(now))
}

func concurrentVer(numCpu int) {
	log.Println("Concurrent version with num cpus:", numCpu)
	now := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(12)
	for i := 0; i < 12; i++ {
		go func(i int) {
			defer wg.Done()
			execIn1s()
		}(i)
	}
	wg.Wait()

	log.Println("since: ", time.Since(now))

}

func sleepIn1s() {
	now := time.Now()
	time.Sleep(time.Second)
	log.Println("time execute a heavy load function: ", time.Since(now))
}
