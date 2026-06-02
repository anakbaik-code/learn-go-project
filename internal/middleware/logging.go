package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader "membajak" fungsi aslinya untuk mencatat status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Inisialisasi wrapper, default status adalah 200 OK
		rw := &responseWriter{w, http.StatusOK}

		// Teruskan ke handler berikutnya
		next.ServeHTTP(rw, r)

		// Eksekusi logging setelah handler selesai (Post-processing)
		slog.Info("incoming_request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", rw.statusCode),
			slog.Duration("latency", time.Since(start)),
			slog.String("ip", r.RemoteAddr),
		)
	})
}
