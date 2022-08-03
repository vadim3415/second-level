package handler

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    start := time.Now()
    next.ServeHTTP(w, req)
    logrus.Printf("method: %s  URI: %s  lead time: %s", req.Method, req.RequestURI, time.Since(start))
  })
}