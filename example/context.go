package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("Server: hello handler started")
	defer fmt.Println("Server: hello handler ended")

	select {
		case <-time.After(4 * time.Second):
			fmt.Fprintf(w, "hello there\n")
		case <-ctx.Done():
			err := ctx.Err()
			fmt.Println("server:", err)
			internalError := http.StatusInternalServerError
			http.Error(w, err.Error(), internalError)

	}
}

func main() {
	http.HandleFunc("/hello", hello)
	port := os.Args[1]
	http.ListenAndServe(":" + port, nil)
}