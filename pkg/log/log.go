package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gobot/pkg/settings"
	"strings"
)

var (
	logger    *zap.Logger
	zapConfig zap.Config
)

func Init() {
	level := strings.ToLower(settings.Settings.Log.Level)
	zapConfig.EncoderConfig = zap.NewProductionEncoderConfig()
	switch level {
	case "debug":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		break
	case "info":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		break
	case "warn":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
		break
	case "error":
		zapConfig.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
		break
	default:
		logger.Fatal("log conf only allow [debug, info, warn, error], please check your configure.")
	}

	logFile := settings.Settings.Log.File
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if logFile == "" {
		zapConfig.OutputPaths = []string{"stderr"}
		zapConfig.ErrorOutputPaths = []string{"stderr"}
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapConfig.Encoding = "console"
	} else {
		zapConfig.Sampling = &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		}
		zapConfig.Encoding = "json"
		zapConfig.OutputPaths = []string{logFile}
		zapConfig.ErrorOutputPaths = []string{logFile}
	}
	logger, _ = zapConfig.Build()
}

func Debug(v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Debug(fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Debug(fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Info(fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Warn(fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Warn(fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Error(fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Error(fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Fatal(fmt.Sprint(v...))
}

func Fatalf(format string, v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Fatal(fmt.Sprintf(format, v...))
}

func Panic(v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Panic(fmt.Sprint(v...))
}

func Panicf(format string, v ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(1)).Panic(fmt.Sprintf(format, v...))
}
