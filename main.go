package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Server is running")
}

func uploadHandler(writer http.ResponseWriter, request *http.Request) {
	file, header, err := parseFormFile(request, "file")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Println("Uploaded file:", header.Filename)

	err = saveFile(header.Filename, file)
	if err != nil {
		http.Error(writer, "Error saving file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(writer, "File uploaded successfully!")
}

func parseFormFile(request *http.Request, name string) (multipart.File, *multipart.FileHeader, error) {
	err := request.ParseMultipartForm(10 << 20)

	if err != nil {
		return nil, nil, err
	}

	file, header, err := request.FormFile(name)
	if err != nil {
		return nil, nil, err
	}

	return file, header, nil
}

func saveFile(fileName string, file io.Reader) error {
	out, err := os.Create("./uploads/" + fileName)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	router := chi.NewRouter()

	router.Get("/", index)
	router.Post("/upload", uploadHandler)

	fmt.Println("Listening on :8000...")

	http.ListenAndServe(":8000", router)
}
