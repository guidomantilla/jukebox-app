package config

import (
	"fmt"
	"jukebox-app/src/misc/environment"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var singletonLogger *zap.Logger

func StopMonitoring() {

	// Stop Zap
	_ = singletonLogger.Sync()

	// Stop Sentry
	sentry.Flush(2 * time.Second)
}

func InitMonitoring(environment environment.Environment) {

	// Setup & Init Sentry
	sentryOptions := sentry.ClientOptions{
		Dsn:         environment.GetValue("SENTRY_DSN").AsString(),
		Environment: environment.GetValue("SENTRY_ENVIRONMENT").AsString(),
		Release:     environment.GetValue("SENTRY_RELEASE").AsString(),
		Debug:       true,
	}

	if err := sentry.Init(sentryOptions); err != nil {
		log.Fatalln(fmt.Sprintf("sentry.Init: %s", err))
	}

	// Setup Zap
	level := zapcore.Level(0)
	if err := level.UnmarshalText([]byte(environment.GetValue("LOG_LEVEL").AsString())); err != nil {
		log.Fatalln(fmt.Sprintf("invalid zap log level: %s", err))
	}

	loggerConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Init Zap with Sentry Hooks for error level logs
	var err error
	if singletonLogger, err = loggerConfig.Build(zap.Hooks(sentryHook)); err != nil {
		log.Fatalln(err)
	}
	zap.ReplaceGlobals(singletonLogger)
}

func sentryHook(entry zapcore.Entry) error {
	if entry.Level == zapcore.ErrorLevel {
		sentry.CaptureException(fmt.Errorf("%s, Line No: %d :: %s", entry.Caller.File, entry.Caller.Line, entry.Message))
	}
	return nil
}
