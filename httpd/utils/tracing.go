package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type key int

const (
	requestIDKey key = 0
)

func Tracing() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func nextRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
