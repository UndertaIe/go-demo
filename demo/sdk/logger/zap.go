package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func DescOfZap() {
	fmt.Println("zap demo")
}

func (l L) BasicOfZap() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()
	url := "https://github.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
	logger.Sync() // flushes buffer, if any
}
