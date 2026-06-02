package logger

import (
	"log"
	"log/slog"
	"os"
)

// setup slog file

func InitLogger() *os.File {
	filepath := "storage/logs/app.log"
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("can't open file log %s : %v", filepath, err)
	}

	handler := slog.NewJSONHandler(file, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return file
}
