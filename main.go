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
	for j := 0; j < 4_000_000_000; j++ {
	}
	log.Println("time execute a heavy load function: ", time.Since(now))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("===================")
	fmt.Println("GO VER: ", runtime.Version())
	procs, _ := strconv.Atoi(os.Getenv("GOMAXPROCS"))
	n := runtime.GOMAXPROCS(procs)
	fmt.Println("num pros:", n)

	writeFile()
	// concurrentVer(n) // expect time: t*[10/n]
	kernelJob()
}

func sequentialVer() {
	log.Println("Sequential version")
	now := time.Now()

	for i := 0; i < 20; i++ {
		execIn1s()
	}

	log.Println("since: ", time.Since(now))
}

func kernelJob() {
	now := time.Now()
	j := os.Getenv("JOBS")
	n, _ := strconv.Atoi(j)
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			execIn1s()
			writeFile()
		}()
	}
	wg.Wait()
	log.Println("since: ", time.Since(now))
}

func writeFile() {
	now := time.Now()
	b := make([]byte, 300000000)
	os.WriteFile(fmt.Sprintf("%d.txt", time.Now().UnixNano()), b, 0644)
	log.Println("Write file exec in: ", time.Since(now))

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
