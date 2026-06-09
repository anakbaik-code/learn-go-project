package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"go-dbsqlc/internal/logger"
	"log/slog"
	"net/http"
	"time"
)

func generateRequestId() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqId := generateRequestId()

		ctx := context.WithValue(r.Context(), logger.RequestIDKey, reqId)
		r = r.WithContext(ctx)
		// Teruskan ke handler berikutnya
		next.ServeHTTP(w, r)

		// Eksekusi logging setelah handler selesai (Post-processing)
		slog.InfoContext(ctx, "incoming_request",
			"method", r.Method,
			"psth", r.URL.Path,
			"latency_ms", time.Since(start).Milliseconds(),
			"ip", r.RemoteAddr,
		)
	})
}
