package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// demo.PartitionRoundRobin()
	// demo.PartitionMurmur2()
	// demo.PartitionCustom()
	// demo.ProduceWithoutBatching()
	// demo.ProduceWithBatching()
	// demo.ProduceWithAck()
	// demo.ConsumerOffset()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		f, _ := os.Open("data.txt")
		defer func() {
			log.Printf("zero copy since %fs", time.Since(now).Seconds())
			f.Close()
		}()
		io.Copy(w, f)
	})

	http.HandleFunc("/copy", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		data, _ := os.ReadFile("data.txt")
		defer func() {
			log.Printf("copy since %fs", time.Since(now).Seconds())
		}()
		w.Write(data)
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println(err)
	}

}
