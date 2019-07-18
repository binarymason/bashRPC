package main

import (
	"log"
	"net/http"
	"os"
)

func initLog(logPath string) (*os.File, error) {
	return os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// logRequest is a closure that allows logging of the request as well as the response
func logRequest(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		lrw := loggingResponseWrapper(w)
		wrappedHandler.ServeHTTP(lrw, req)
		statusCode := lrw.statusCode
		log.Println(req.RemoteAddr, req.Method, req.URL, statusCode, http.StatusText(statusCode))
	})
}

func loggingResponseWrapper(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
