package internal

//go:generate mockgen -destination=../internal/mocks/logging.go -package=mocks -source=../internal/logging.go

import (
	"context"
	"net/http"
	"time"
)

// RequestLog is the interface for logging at the request level where a context should always be present
type RequestLog interface {
	ContextLog

	// ForContext creates a log for the specific context so the context doesn't have to be passed to each log method
	ForContext(ctx context.Context) Log // log is returned because the context is already known
}

// BackgroundLog is the interface for background type task logging that is not directly related to a request. It
// still supports context based logging when a context is available.
type BackgroundLog interface {
	Log
	ContextLog

	// ChildLog will create a named child log below the current one with its own name so entries for it can be easily
	// separated.  Names follow a . notation ie parent.child
	ChildLog(name string) BackgroundLog
}

// StartupLog is the interface for the startup log which is meant to log startup and shutdown activity.
type StartupLog interface {
	Log
	ContextLog
}

// log is a base interface not directly injectable but embedded in other logging interfaces to reuse method signatures
type Log interface {
	// Debug logs a message at debug level
	Debug(args ...interface{})

	// Debugf logs a formatted message at debug level
	Debugf(template string, args ...interface{})

	// Debugw logs a message at debug level along with a set of key value pairs that are turned into structured logging fields
	Debugw(message string, keyValues ...interface{})

	// Info logs a message at info level
	Info(args ...interface{})

	// Infof logs a formatted message at info level
	Infof(template string, args ...interface{})

	// Infow logs a message at info level along with a set of key value pairs that are turned into structured logging fields
	Infow(message string, keyValues ...interface{})

	// Warn logs a message at warn level
	Warn(args ...interface{})

	// Warnf logs a formatted message at warn level
	Warnf(template string, args ...interface{})

	// Warnw logs a message at warn level along with a set of key value pairs that are turned into structured logging fields
	Warnw(message string, keyValues ...interface{})

	// Error logs a message at error level
	Error(args ...interface{})

	// Errorf logs a formatted message at error level
	Errorf(template string, args ...interface{})

	// Errorw logs a message at error level along with a set of key value pairs that are turned into structured logging fields
	Errorw(message string, keyValues ...interface{})

	// Err logs a Go error with a message at the error level
	Err(err error, message string)

	// Errf logs a Go error with a formatted message at the error level
	Errf(err error, template string, args ...interface{})

	// Errw logs a Go error at the error level along with a set of key value pairs that are turned into structured logging fields
	Errw(err error, message string, keyValues ...interface{})
}

// ContextLog is a base interface that contains context sensitive logging methods so the method signatures can be reused
type ContextLog interface {
	// DebugCtx logs a message at debug level
	DebugCtx(ctx context.Context, message string)

	// DebugfCtx logs a formatted message at debug level
	DebugfCtx(ctx context.Context, template string, args ...interface{})

	// DebugwCtx logs a message at debug level along with a set of key value pairs that are turned into structured logging fields
	DebugwCtx(ctx context.Context, message string, keyValues ...interface{})

	// InfoCtx logs a message at info level
	InfoCtx(ctx context.Context, message string)

	// InfofCtx logs a formatted message at info level
	InfofCtx(ctx context.Context, template string, args ...interface{})

	// InfowCtx logs a message at info level along with a set of key value pairs that are turned into structured logging fields
	InfowCtx(ctx context.Context, message string, keyValues ...interface{})

	// WarnCtx logs a message at warn level
	WarnCtx(ctx context.Context, message string)

	// WarnfCtx logs a formatted message at warn level
	WarnfCtx(ctx context.Context, template string, args ...interface{})

	// WarnwCtx logs a message at warn level along with a set of key value pairs that are turned into structured logging fields
	WarnwCtx(ctx context.Context, message string, keyValues ...interface{})

	// ErrorCtx logs a message at error level
	ErrorCtx(ctx context.Context, message string)

	// ErrorfCtx logs a formatted message at error level
	ErrorfCtx(ctx context.Context, template string, args ...interface{})

	// ErrorwCtx logs a message at error level along with a set of key value pairs that are turned into structured logging fields
	ErrorwCtx(ctx context.Context, message string, keyValues ...interface{})

	// ErrCtx logs a Go error with a message at the error level
	ErrCtx(ctx context.Context, err error, message string)

	// ErrfCtx logs a Go error with a formatted message at the error level
	ErrfCtx(ctx context.Context, err error, template string, args ...interface{})

	// ErrwCtx logs a Go error at the error level along with a set of key value pairs that are turned into structured logging fields
	ErrwCtx(ctx context.Context, err error, message string, keyValues ...interface{})
}

// IncomingHttpRequestSampler is an interface implemented to control the sampling of incoming http request logs and
// when they actually should be written.
type IncomingHttpRequestSampler interface {
	// ShouldLogStart indicates whether the start of the request should be logged
	ShouldLogStart(r *http.Request) bool

	// ShouldLogEnd determines whether the end of the request should be logged
	ShouldLogEnd(r *http.Request, responseCode int, duration time.Duration) bool
}

