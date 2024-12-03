package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

var (
	// Logger is the global variable for logging
	Logger *zap.SugaredLogger
	DEBUG  = os.Getenv("DEBUG") // Set DEBUG to true
)

const (
	SamplingConfigInitial    = 100
	SamplingConfigThereafter = 100
	// Log Levels
	ErrorLevel   = zap.ErrorLevel
	WarningLevel = zap.WarnLevel
	InfoLevel    = zap.InfoLevel
	DebugLevel   = zap.DebugLevel
)

// Options Logger options
type Options struct {
	Level zapcore.Level
	Core  *zapcore.Core
}

// OptFunc used for configuring to set NewLogger Options
type OptFunc func(*Options)

// defaultOpts Set default for struct Options
func defaultOpts() Options {
	return Options{
		Level: InfoLevel,
	}
}

// WithCore set Options core
func WithCore(core zapcore.Core) OptFunc {
	return func(o *Options) {
		o.Core = &core
	}
}

// WithLevel set options level
func WithLevel(level zapcore.Level) OptFunc {
	return func(o *Options) {
		o.Level = level
	}
}

// init a go internal function that runs once package log is imported; the more you know ;)
func init() {
	// Set the default logger
	Logger = NewLogger(func(o *Options) {
		// If environmental variable DEBUG equal true then log at debug level
		if DEBUG == "true" {
			o.Level = DebugLevel
		}
	})
}

// InitializeLogger initialize the Logger variable
func InitializeLogger(optFns ...OptFunc) {
	Logger = NewLogger(optFns...)
}

// NewLogger create a new logger
func NewLogger(optFns ...OptFunc) *zap.SugaredLogger {
	options := defaultOpts()
	for _, fn := range optFns {
		fn(&options)
	}

	// If options.Core exist the function will end and return this new logger.
	if options.Core != nil {
		logger := zap.New(*options.Core)
		return logger.Sugar()
	}

	atom := zap.NewAtomicLevelAt(options.Level)
	encoderConfig := zap.NewProductionEncoderConfig()

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		atom,
	))

	return logger.Sugar()
}

// Debug logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Debug(msg string, keysAndValues ...interface{}) {
	Logger.Debugw(msg, keysAndValues...)
}

// Info logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Info(msg string, keysAndValues ...interface{}) {
	Logger.Infow(msg, keysAndValues...)
}

// Warn logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Warn(msg string, keysAndValues ...interface{}) {
	Logger.Warnw(msg, keysAndValues...)
}

// Error logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Error(msg string, keysAndValues ...interface{}) {
	Logger.Errorw(msg, keysAndValues...)
}

// Fatal logs a message with some additional context, then calls os.Exit. The variadic key-value pairs are treated as they are in With.
func Fatal(msg string, keysAndValues ...interface{}) {
	Logger.Fatalw(msg, keysAndValues...)
}

// CapturesLogs capture the logs
func CapturesLogs(level zapcore.Level) *observer.ObservedLogs {
	observerCore, observedlogs := observer.New(level)
	InitializeLogger(WithCore(observerCore))

	return observedlogs
}

// String wraps the zap.String function
func String(key, val string) zap.Field {
	if key == "" {
		return zap.Skip()
	}
	return zap.String(key, val)
}

// Int wraps the zap.Int function
func Int(key string, val int) zap.Field {
	if key == "" {
		return zap.Skip()
	}
	return zap.Int(key, val)
}

// Bool wraps the zap.Bool function
func Bool(key string, val bool) zap.Field {
	if key == "" {
		return zap.Skip()
	}
	return zap.Bool(key, val)
}

// Any wraps the zap.Any function
func Any(key string, val interface{}) zap.Field {
	if key == "" {
		return zap.Skip()
	}

	return zap.Any(key, val)
}
