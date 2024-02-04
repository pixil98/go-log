package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

const (
	loggerKey key = "logger"
)

type key string

func GetLogger(ctx context.Context) *logrus.Logger {
	val := ctx.Value(loggerKey)
	if val == nil {
		return NewLogger()
	}

	logger, ok := val.(*logrus.Logger)
	if !ok {
		return NewLogger()
	}

	return logger
}

func SetLogger(ctx context.Context, logger logrus.FieldLogger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
