package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Server is running")
	})

	fmt.Println("Listening on :8000...")

	http.ListenAndServe(":8000", nil)
}
