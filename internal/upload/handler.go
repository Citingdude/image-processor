package upload

import (
	"fmt"
	"net/http"
)

func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	file, header, err := ParseFormFile(request, "file")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Println("Uploaded file:", header.Filename)

	err = SaveFile(header.Filename, file)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(writer, "File uploaded successfully!")
}
