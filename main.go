package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
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
		byteDump, _ := httputil.DumpRequest(r, true)
		fmt.Fprintf(w, "Service " + count)
		fmt.Fprintf(w, "\n" + string(byteDump))
	})
	http.ListenAndServe(":" + port, nil)
}
