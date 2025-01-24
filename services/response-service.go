package services

import "net/http"

// Create a custom response writer to capture the output
type ResWriter struct {
	http.ResponseWriter
	Body []byte
}

func (rw *ResWriter) Write(b []byte) (int, error) {
	rw.Body = append(rw.Body, b...)
	return rw.ResponseWriter.Write(b)
}
