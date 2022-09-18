package config

import (
	"fmt"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"jukebox-app/pkg/environment"
	"jukebox-app/pkg/properties"
)

const (
	CMD_PROPERTY_SOURCE_NAME = "CMD"
	OS_PROPERTY_SOURCE_NAME  = "OS"
)

var (
	_singletonEnvironment environment.Environment
	_singletonLogger      *zap.Logger
)

func InitConfig(cmdArgs *[]string) environment.Environment {

	zap.L().Info("server starting up - setting up configuration & logging")

	// Load CMD and OS variables
	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OS_PROPERTY_SOURCE_NAME,
		properties.NewDefaultProperties(properties.FromArray(&osArgs)))

	cmdSource := properties.NewDefaultPropertySource(CMD_PROPERTY_SOURCE_NAME,
		properties.NewDefaultProperties(properties.FromArray(cmdArgs)))

	_singletonEnvironment = environment.NewDefaultEnvironment(environment.WithPropertySources(osSource, cmdSource))

	// Setup & Init Sentry
	sentryOptions := sentry.ClientOptions{
		Dsn:         _singletonEnvironment.GetValue("SENTRY_DSN").AsString(),
		Environment: _singletonEnvironment.GetValue("SENTRY_ENVIRONMENT").AsString(),
		Release:     _singletonEnvironment.GetValue("SENTRY_RELEASE").AsString(),
		Debug:       true,
	}

	var err error
	if err = sentry.Init(sentryOptions); err != nil {
		zap.L().Fatal(fmt.Sprintf("server starting up - error setting up sentry: %s", err.Error()))
	}

	// Setup Zap
	level := zapcore.Level(0)
	if err = level.UnmarshalText([]byte(_singletonEnvironment.GetValue("LOG_LEVEL").AsString())); err != nil {
		zap.L().Fatal(fmt.Sprintf("server starting up - invalid zap log level: %s", err.Error()))
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
		zap.L().Fatal(fmt.Sprintf("server starting up - error hooking up sentry and zap: %s", err.Error()))
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
	if entry.Level == zapcore.ErrorLevel || entry.Level == zapcore.FatalLevel {
		sentry.CaptureException(fmt.Errorf("%s, Line No: %d :: %s", entry.Caller.File, entry.Caller.Line, entry.Message))
	}
	return nil
}

func StopConfig() error {

	zap.L().Info("server shutting down - stopping configuration & logging")

	// Stop Zap
	_ = _singletonLogger.Sync()

	// Stop Sentry
	sentry.Flush(2 * time.Second)

	zap.L().Info("server shutting down - configuration & logging stopped")
	return nil
}
