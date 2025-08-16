package upload

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func ParseFormFile(request *http.Request, name string) (multipart.File, *multipart.FileHeader, error) {
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

func SaveFile(fileName string, file io.Reader) error {
	err := os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		return err
	}

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
