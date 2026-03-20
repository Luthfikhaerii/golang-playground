package logger

import (
	"os"

	"go.uber.org/zap"
)

var Log *zap.Logger

func Init() {
	env := os.Getenv("APP_ENV")

	if env == "production" {
		Log, _ = zap.NewProduction()
	} else {
		Log, _ = zap.NewDevelopment()
	}

	// optional: tambah field global
	// Log = Log.With(
	// 	zap.String("service", "golang-playground"),
	// )
}
