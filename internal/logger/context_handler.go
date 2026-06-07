package logger

import (
	"context"
	"log/slog"
)

type ctxKey string
const RequestIDKey ctxKey = "request_id"

type ContextHandler struct {
	slog.Handler
}

// Handle Context 
func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if ctx != nil {
		if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
			// Otomatis tambahkan ke atribut log tanpa perlu diketik manual di service
			r.AddAttrs(slog.String("request_id", reqID))
		}
	}
	return h.Handler.Handle(ctx, r)
}