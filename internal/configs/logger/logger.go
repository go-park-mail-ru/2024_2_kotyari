package logger

import (
	"github.com/natefinch/lumberjack"
	"log/slog"
	"os"
)

// InitLogger инициализирует логгер для разных режимов (dev/prod/local)
func InitLogger() *slog.Logger {
	var (
		logger *slog.Logger
		env    = os.Getenv("APP_ENV")
	)

	switch env {
	case "local":
		logger = newLogger(slog.LevelDebug)
	case "dev":
		logger = newLogger(slog.LevelInfo)
	default:
		logFile := &lumberjack.Logger{
			Filename:   "app.log",
			MaxSize:    10, // MB
			MaxBackups: 3,  // Сколько старых файлов хранить
			MaxAge:     3,  // Сколько дней хранить файлы
			Compress:   true,
		}

		logger = slog.New(slog.NewJSONHandler(
			logFile,
			&slog.HandlerOptions{
				Level: slog.LevelError,
			}))
	}

	return logger
}
