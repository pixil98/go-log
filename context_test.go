package log

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetLogger(t *testing.T) {
	testLogger := logrus.New()
	testEntry := testLogger.WithField("test", "value")

	tests := map[string]struct {
		ctx      context.Context
		expected logrus.FieldLogger
	}{
		"returns new logger when context has no logger": {
			ctx:      context.Background(),
			expected: nil,
		},
		"returns logger when context has *logrus.Logger": {
			ctx:      context.WithValue(context.Background(), loggerKey, testLogger),
			expected: testLogger,
		},
		"returns entry when context has *logrus.Entry": {
			ctx:      context.WithValue(context.Background(), loggerKey, testEntry),
			expected: testEntry,
		},
		"returns new logger when context has wrong type": {
			ctx:      context.WithValue(context.Background(), loggerKey, "not a logger"),
			expected: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := GetLogger(tc.ctx)

			if tc.expected == nil {
				if result == nil {
					t.Error("expected non-nil logger, got nil")
				}
			} else if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestSetLogger(t *testing.T) {
	testLogger := logrus.New()
	testEntry := testLogger.WithField("test", "value")

	tests := map[string]struct {
		logger logrus.FieldLogger
	}{
		"stores *logrus.Logger": {
			logger: testLogger,
		},
		"stores *logrus.Entry": {
			logger: testEntry,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := SetLogger(context.Background(), tc.logger)
			result := GetLogger(ctx)

			if result != tc.logger {
				t.Errorf("expected %v, got %v", tc.logger, result)
			}
		})
	}
}
