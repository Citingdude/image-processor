# Go Image Processor 🖼️⚡

A simple HTTP service written in Go that accepts an uploaded image, processes it (e.g., grayscale, pixelate, add text), and returns the transformed image.

This is a learning project to explore:
- Go's `net/http` server
- File uploads & multipart forms
- Image decoding & encoding
- Basic image manipulation
- Streaming binary data in HTTP responses

---

## 🚀 Features

- **Upload images** via HTTP POST
- **Process images** (start with grayscale, then experiment with other effects)
- **Return processed image** directly as a response
- Minimal dependencies, mostly standard library

---

## 📦 Getting Started

### 1. Clone & Init
```bash
git clone https://github.com/yourusername/go-image-processor.git
cd go-image-processor
go mod init image-processor
```

### 2. Run the server
```bash
go run main.go
```

By default, it will listen on `http://localhost:8000`.

## Example Usage

### Upload an image with curl

```bash
curl -X POST -F "file=@your-image.jpg" http://localhost:8080/upload --output processed.jpg
```
- file — the form field name for the uploaded image
- The processed image will be saved as processed.jpg