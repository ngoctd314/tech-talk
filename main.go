package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func execIn1s() {
	now := time.Now()
	for j := 0; j < 10_000_000_000; j++ {
	}
	log.Println("time execute a heavy load function: ", time.Since(now))
}

func execIn2s() {
	now := time.Now()
	for j := 0; j < 4_000_000_000; j++ {
	}
	log.Println("time execute a heavy load function: ", time.Since(now))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("===================")
	fmt.Println("GO VER: ", runtime.Version())
	numCpu := os.Getenv("GOMAXPROCS")
	procs, _ := strconv.Atoi(numCpu)
	n := runtime.GOMAXPROCS(procs)
	fmt.Println("num pros:", n)

	// sequentialVer()
	concurrentVer(n) // expect time: t*[10/n]
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

	v := os.Getenv("JOBS")
	n, _ := strconv.Atoi(v)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			execIn1s()
		}(i)
	}
	fmt.Println("num goroutines: ", runtime.NumGoroutine())
	wg.Wait()

	log.Println("since: ", time.Since(now))

}

func sleepIn1s() {
	now := time.Now()
	time.Sleep(time.Second)
	log.Println("time execute a heavy load function: ", time.Since(now))
}
