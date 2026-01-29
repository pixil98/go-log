package log

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewLogger(t *testing.T) {
	tests := map[string]struct {
		opts          []LoggerOpt
		expectedLevel logrus.Level
	}{
		"returns logger with default level when no options": {
			opts:          nil,
			expectedLevel: logrus.InfoLevel,
		},
		"applies WithLevel option": {
			opts:          []LoggerOpt{WithLevel(logrus.DebugLevel)},
			expectedLevel: logrus.DebugLevel,
		},
		"applies multiple options": {
			opts: []LoggerOpt{
				WithLevel(logrus.WarnLevel),
				WithLevel(logrus.ErrorLevel),
			},
			expectedLevel: logrus.ErrorLevel,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			logger := NewLogger(tc.opts...)

			if logger == nil {
				t.Fatal("expected non-nil logger")
			}

			if logger.GetLevel() != tc.expectedLevel {
				t.Errorf("expected level %v, got %v", tc.expectedLevel, logger.GetLevel())
			}
		})
	}
}

func TestWithLevel(t *testing.T) {
	tests := map[string]struct {
		level logrus.Level
	}{
		"sets debug level": {
			level: logrus.DebugLevel,
		},
		"sets info level": {
			level: logrus.InfoLevel,
		},
		"sets warn level": {
			level: logrus.WarnLevel,
		},
		"sets error level": {
			level: logrus.ErrorLevel,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			logger := logrus.New()
			opt := WithLevel(tc.level)
			opt(logger)

			if logger.GetLevel() != tc.level {
				t.Errorf("expected level %v, got %v", tc.level, logger.GetLevel())
			}
		})
	}
}
