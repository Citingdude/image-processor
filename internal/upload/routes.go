package upload

import "github.com/go-chi/chi/v5"

func Routes(r chi.Router) {
	r.Route("/upload", func(r chi.Router) {
		r.Post("/", UploadHandler)
	})
}
