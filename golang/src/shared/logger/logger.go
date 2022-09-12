package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Setup() {
	var err error

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableStacktrace: true,
		Encoding:          "console",
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:      "severity",
			NameKey:       "logger",
			CallerKey:     "caller",
			StacktraceKey: "stack_trace",
			TimeKey:       "time",
			MessageKey:    "message",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeTime:    zapcore.RFC3339NanoTimeEncoder,
			EncodeLevel:   zapcore.CapitalColorLevelEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Info(message string, args ...interface{}) {
	log.Sugar().Infof("%s %s", message, args)
}

func Debug(message string, args ...interface{}) {
	log.Sugar().Debugf("%s %s", message, args)
}

func Error(message string, args ...interface{}) {
	log.Sugar().Errorf("%s %s", message, args)
}
