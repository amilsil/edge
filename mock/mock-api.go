package main

import (
	"fmt"
	"net/http"
	"strings"
)

func ping(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message

	fmt.Println("Responding with ", message)
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", ping)
	fmt.Println("Starting Mock Api")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
