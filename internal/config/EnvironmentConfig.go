package config

import (
	"fmt"
	"jukebox-app/pkg/environment"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _singletonEnvironment environment.Environment
var _singletonLogger *zap.Logger

func StopConfig() {
	// Stop Zap
	_ = _singletonLogger.Sync()

	// Stop Sentry
	sentry.Flush(2 * time.Second)
}

func InitConfig(cmdArgs *[]string) environment.Environment {

	// Load CMD and OS variables
	_singletonEnvironment = environment.LoadEnvironment(cmdArgs)

	// Setup & Init Sentry
	sentryOptions := sentry.ClientOptions{
		Dsn:         _singletonEnvironment.GetValue("SENTRY_DSN").AsString(),
		Environment: _singletonEnvironment.GetValue("SENTRY_ENVIRONMENT").AsString(),
		Release:     _singletonEnvironment.GetValue("SENTRY_RELEASE").AsString(),
		Debug:       true,
	}

	var err error
	if err = sentry.Init(sentryOptions); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Setup Zap
	level := zapcore.Level(0)
	if err = level.UnmarshalText([]byte(_singletonEnvironment.GetValue("LOG_LEVEL").AsString())); err != nil {
		log.Fatalf("invalid zap log level: %s", err)
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
	if _singletonLogger, err = loggerConfig.Build(zap.Hooks(sentryHook)); err != nil {
		log.Fatalln(err)
	}
	zap.ReplaceGlobals(_singletonLogger)

	if loggerConfig.Level.String() == "debug" {
		for _, source := range _singletonEnvironment.GetPropertySources() {
			sourceMap := source.AsMap()
			name, internalMap := sourceMap["name"], sourceMap["value"].(map[string]string)
			for key, value := range internalMap {
				zap.L().Debug(fmt.Sprintf("source name: %s, key: %s, value: %s", name, key, value))
			}
		}
	}

	return _singletonEnvironment
}

func sentryHook(entry zapcore.Entry) error {
	if entry.Level == zapcore.ErrorLevel {
		sentry.CaptureException(fmt.Errorf("%s, Line No: %d :: %s", entry.Caller.File, entry.Caller.Line, entry.Message))
	}
	return nil
}
