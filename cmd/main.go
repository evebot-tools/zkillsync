package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"time"

	"github.com/evebot-tools/utils"
	"github.com/evebot-tools/zkillsync/internal"
	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              utils.GetEnv("SENTRY_DSN", ""),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	defer sentry.Flush(5 * time.Second)
	internal.Run()
}
