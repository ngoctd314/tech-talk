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

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("===================")
	fmt.Println("GO VER: ", runtime.Version())
	procs, _ := strconv.Atoi(os.Getenv("GOMAXPROCS"))
	n := runtime.GOMAXPROCS(procs)
	fmt.Println("num pros:", n)

	ver, _ := strconv.Atoi(os.Getenv("VER"))
	if ver == 0 {
		sequentialVer()
	} else {
		concurrentVer() // expect time: t*[10/n]
	}
}

func concurrentVer() {
	now := time.Now()

	wg := sync.WaitGroup{}
	jobs, _ := strconv.Atoi(os.Getenv("JOBS"))
	wg.Add(jobs)
	for i := 0; i < jobs; i++ {
		go func(i int) {
			defer wg.Done()
			execIn1s()
		}(i)
	}

	wg.Wait()

	log.Println("since: ", time.Since(now))
}
func execIn2s() {
	now := time.Now()
	for j := 0; j < 8_000_000_000; j++ {
	}
	log.Println("time execute a heavy load function: ", time.Since(now))
}
func execIn1s() {
	now := time.Now()
	for j := 0; j < 4_000_000_000; j++ {
	}
	log.Println("time execute a heavy load function: ", time.Since(now))
}

// func sleepIn1s() {
// 	now := time.Now()
// 	time.Sleep(time.Second)
// 	log.Println("time execute a heavy load function: ", time.Since(now))
// }

// func writeFile() {
// 	now := time.Now()
// 	b := make([]byte, 300000000)
// 	os.WriteFile(fmt.Sprintf("%d.txt", time.Now().UnixNano()), b, 0644)
// 	log.Println("Write file exec in: ", time.Since(now))

// }

// func kernelJob() {
// 	now := time.Now()
// 	j := os.Getenv("JOBS")
// 	n, _ := strconv.Atoi(j)
// 	wg := sync.WaitGroup{}
// 	wg.Add(n)
// 	for i := 0; i < n; i++ {
// 		go func() {
// 			defer wg.Done()
// 			execIn1s()
// 			writeFile()
// 		}()
// 	}
// 	wg.Wait()
// 	log.Println("since: ", time.Since(now))
// }

func sequentialVer() {
	log.Println("Sequential version")
	now := time.Now()

	jobs, _ := strconv.Atoi(os.Getenv("JOBS"))

	for i := 0; i < jobs; i++ {
		execIn1s()
	}

	log.Println("since: ", time.Since(now))
}

// func execIn2s() {
// 	now := time.Now()
// 	for j := 0; j < 16_000_000_000; j++ {
// 	}
// 	log.Println("time execute a heavy load function: ", time.Since(now))
// }
