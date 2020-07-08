package main

import (
	"fmt"
	"net/http"
	"syscall"
)

const DefaultPort = "8080"

func main() {
        count := "ONE"
	if sc, ok := syscall.Getenv("SIMPLE_COUNT"); ok {
		count = sc
	}

	port := DefaultPort
	if sp, ok := syscall.Getenv("SIMPLE_PORT"); ok {
		port = sp
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service number " + count)
	})
	http.ListenAndServe(":" + port, nil)
}
