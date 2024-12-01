package log_test

import (
	"testing"

	"aoc2024/pkg/log"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

// TestInitializeLogger tests for function InitializeLogger
func TestInitializeLogger(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Test InitializeLogger Success": testInitializeLoggerSuccess,
		"Test InitializeLogger Level":   testInitializeLoggerLevel,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

// TestString tests for function String
func TestString(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Test String Success":     testStringSuccess,
		"Test String Missing Key": testStringMissingKey,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

// TestInt tests for function Int
func TestInt(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Test String Success":     testIntSuccess,
		"Test String Missing Key": testIntMissingKey,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

// Setup functions for capturing logs
func Setup() *observer.ObservedLogs {
	observerCore, observedLogs := observer.New(zap.DebugLevel)
	log.InitializeLogger(log.WithCore(observerCore))
	return observedLogs
}

func testInitializeLoggerSuccess(t *testing.T) {
	ptrPreviousLogger := log.Logger

	log.InitializeLogger()

	assert.NotEqual(t, ptrPreviousLogger, log.Logger,
		"InitializeLogger should assign a new log.Logger is still references the old pointer",
	)
}

func testInitializeLoggerLevel(t *testing.T) {
	for _, test := range []struct {
		description string
		level       zapcore.Level
	}{
		{
			description: "Debug Level",
			level:       log.DebugLevel,
		},
		{
			description: "Info Level",
			level:       log.InfoLevel,
		},
		{
			description: "Warning Level",
			level:       log.WarningLevel,
		},
		{
			description: "Error Level",
			level:       log.ErrorLevel,
		},
	} {
		t.Run(test.description, func(t *testing.T) {
			log.InitializeLogger(log.WithLevel(test.level))
			assert.Equal(t, test.level, log.Logger.Level(), "Expected log Level to reflect the change")
		})
	}
}

func testStringSuccess(t *testing.T) {
	observedLogger := Setup()
	for _, test := range []struct {
		description string
		key         string
		value       string
	}{
		{
			description: "Simple hello world",
			key:         "hello",
			value:       "world",
		},
		{
			description: "Empty value",
			key:         "Okay",
			value:       "",
		},
	} {
		t.Run(test.description, func(t *testing.T) {
			observedLogger.TakeAll() // Empty previous logged entries

			log.Logger.Infow("TEST", log.String(test.key, test.value))

			require.Equal(t, 1, observedLogger.Len(), "Expected number of logged messages do not match")

			logs := observedLogger.TakeAll()

			assert.ElementsMatch(t, []zap.Field{
				zap.String(test.key, test.value),
			}, logs[0].Context)
		})
	}
}

func testStringMissingKey(t *testing.T) {
	observedLogger := Setup()

	log.Logger.Infow("TEST", log.String("", "hello!"))

	require.Equal(t, 1, observedLogger.Len(), "Expected number of logged messages do not match")

	logs := observedLogger.TakeAll()

	assert.ElementsMatch(t, []zap.Field{
		zap.Skip(),
	}, logs[0].Context)
}

func testIntSuccess(t *testing.T) {
	observedLogger := Setup()
	for _, test := range []struct {
		description string
		key         string
		value       int
	}{
		{
			description: "Zero",
			key:         "zero",
			value:       0,
		},
		{
			description: "Eleven",
			key:         "Important-Key",
			value:       11,
		},
	} {
		t.Run(test.description, func(t *testing.T) {
			observedLogger.TakeAll() // Empty previous logged entries

			log.Logger.Infow("TEST", log.Int(test.key, test.value))

			require.Equal(t, 1, observedLogger.Len(), "Expected number of logged messages do not match")

			logs := observedLogger.TakeAll()

			assert.ElementsMatch(t, []zap.Field{
				zap.Int(test.key, test.value),
			}, logs[0].Context)
		})
	}
}

func testIntMissingKey(t *testing.T) {
	observedLogger := Setup()

	log.Logger.Infow("TEST", log.Int("", 10))

	require.Equal(t, 1, observedLogger.Len(), "Expected number of logged messages do not match")

	logs := observedLogger.TakeAll()

	assert.ElementsMatch(t, []zap.Field{
		zap.Skip(),
	}, logs[0].Context)
}
