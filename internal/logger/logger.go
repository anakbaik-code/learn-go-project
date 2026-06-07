package logger

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func InitLogger(appEnv string) {
	programLevel := new(slog.LevelVar)

	var handler slog.Handler

	if strings.ToLower(strings.TrimSpace(appEnv)) == "production" {
		programLevel.Set(slog.LevelInfo)

		logDir := filepath.Join("storage", "logs")
		_ = os.MkdirAll(logDir, 0755)
		logFile := filepath.Join(logDir, "app.log")

		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		var handlerlog slog.Handler
		if err != nil {
			// Jika gagal buka file, masukkan ke handlerlog dengan output Stdout
			handlerlog = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource: true,
				Level:     programLevel,
			})
		} else {
			// Jika sukses, masukkan ke handlerlog dengan output File
			handlerlog = slog.NewJSONHandler(file, &slog.HandlerOptions{
				AddSource: true,
				Level:     programLevel,
			})
		}

		// Masukkan hasil JSON tadi ke dalam ContextHandler kita
		handler = &ContextHandler{Handler: handlerlog}

	} else {
		programLevel.Set(slog.LevelDebug)

		handlerlog := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     programLevel, // slog akan memunculkan debug karena levelnya di-set ke LevelDebug
		})
		handler = &ContextHandler{Handler: handlerlog}

	}

	slog.SetDefault(slog.New(handler))
}
