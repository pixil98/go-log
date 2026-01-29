package log

import (
	"github.com/sirupsen/logrus"
)

type LoggerOpt func(*logrus.Logger)

func WithLevel(level logrus.Level) LoggerOpt {
	return func(l *logrus.Logger) {
		l.SetLevel(level)
	}
}

func NewLogger(opts ...LoggerOpt) *logrus.Logger {
	l := logrus.New()

	for _, opt := range opts {
		opt(l)
	}

	return l
}
