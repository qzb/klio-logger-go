package logger

import (
	"io"
)

// Level type.
type Level string

const (
	// FatalLevel level. Errors causing a command to exit immediately.
	FatalLevel Level = "fatal"
	// ErrorLevel level. Errors which cause a command to fail, but not immediately.
	ErrorLevel Level = "error"
	// WarnLevel level. Information about unexpected situations and minor errors (not causing a command to fail).
	WarnLevel Level = "warn"
	// InfoLevel level. Generally useful information (things happen).
	InfoLevel Level = "info"
	// VerboseLevel level. More granular but still useful information.
	VerboseLevel Level = "verbose"
	// DebugLevel level. Information helpful for command developers.
	DebugLevel Level = "debug"
	// SpamLevel level. Give me EVERYTHING.
	SpamLevel Level = "spam"
)

// Logger.
type Logger struct{}

// New creates new instance of Logger.
func New(output *io.Writer) *Logger

// Tags returns tags used by a logger. ags are prepended to each line produced by a logger.
func (l *Logger) Tags() []string

// Level returns log level used by a logger.
func (l *Logger) Level() Level

// WithLevel creates new logger instance logging at specified level.
func (l *Logger) WithLevel(Level) *Logger

// WithTags creates new logger instance with specified tags. Tags are prepended to each line produced by a logger.
func (l *Logger) WithTags(tags ...string) *Logger

// Printf writes log line. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) *Logger

// Printf writes log line. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) *Logger

// StandardLogger returns logger instance writing to the stdout.
func StandardLogger() *Logger

// ErrorLogger returns logger instance writing to the stderr.
func ErrorLogger() *Logger

// Spam writes a message at level Spam on the standard logger. Arguments are handled in the manner of fmt.Print.
func Spam(v ...interface{}) {}

// Debug writes a message at level Debug on the standard logger. Arguments are handled in the manner of fmt.Print.
func Debug(v ...interface{}) {}

// Verbose writes a message at level Verbose on the standard logger. Arguments are handled in the manner of fmt.Print.
func Verbose(v ...interface{}) {}

// Info writes a message at level Info on the standard logger. Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {}

// Warn writes a message at level Warn on the standard logger. Arguments are handled in the manner of fmt.Print.
func Warn(v ...interface{}) {}

// Error writes a message at level Error on the standard logger. Arguments are handled in the manner of fmt.Print.
func Error(v ...interface{}) {}

// Fatal writes a message at level Fatal on the standard logger. Arguments are handled in the manner of fmt.Print.
func Fatal(v ...interface{}) {}

// Spamf writes a message at level Spam on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Spamf(format string, v ...interface{}) {}

// Debugf writes a message at level Debug on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {}

// Verbosef writes a message at level Verbose on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Verbosef(format string, v ...interface{}) {}

// Infof writes a message at level Info on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {}

// Warnf writes a message at level Warn on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {}

// Errorf writes a message at level Error on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {}

// Fatalf writes a message at level Fatal on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, v ...interface{}) {}
