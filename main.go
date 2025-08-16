package main

import (
	"fmt"
	"image-processor/internal/upload"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	upload.Routes(router)

	fmt.Println("Listening on :8000...")

	http.ListenAndServe(":8000", router)
}
