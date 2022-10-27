package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := make([]byte, 10000000)
		os.WriteFile(fmt.Sprintf("%d.txt", time.Now().UnixNano()), b, 0644)
		log.Println("GET: /")
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", nil)
}
