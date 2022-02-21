// This logger is meant to be used for building Klio commands (https://github.com/g2a-com/klio).
// It writes logs decorated with control sequences interpreted by Klio (https://github.com/g2a-com/klio/blob/main/docs/output-handling.md).
// It doesn't filter or modify messages besides that.

package logger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
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
	// DefaultLevel is an alias for "info" level.
	DefaultLevel = InfoLevel
)

var (
	standardLogger = New(os.Stdout)
	errorLogger    = New(os.Stderr).WithLevel(ErrorLevel)
	levelsMap      = map[string]Level{
		string(FatalLevel):   FatalLevel,
		string(ErrorLevel):   ErrorLevel,
		string(WarnLevel):    WarnLevel,
		string(InfoLevel):    InfoLevel,
		string(VerboseLevel): VerboseLevel,
		string(DebugLevel):   DebugLevel,
		string(SpamLevel):    SpamLevel,
	}
)

// ParseLevel converts level name to Level. It is case insensitive, returns DefaultLevel if value cannot be converted.
func ParseLevel(s string) (level Level, ok bool) {
	level, ok = levelsMap[strings.ToLower(s)]
	if !ok {
		level = DefaultLevel
	}
	return level, ok
}

// Logger.
type Logger struct {
	output     io.Writer
	tags       []string
	level      Level
	linePrefix string
}

// New creates new instance of Logger.
func New(output io.Writer) *Logger {
	l := &Logger{
		output: output,
		tags:   []string{},
		level:  DefaultLevel,
	}

	l.updateLinePrefix()

	return l
}

func (l *Logger) updateLinePrefix() {
	level, err := json.Marshal(l.level)
	if err != nil {
		level = []byte("\"" + DefaultLevel + "\"")
	}
	tags, err := json.Marshal(l.tags)
	if err != nil || string(tags) == "null" {
		tags = []byte("[]")
	}
	l.linePrefix = fmt.Sprintf(
		"\033_klio_log_level %s\033\\\033_klio_tags %s\033\\", level, tags,
	)
}

// Tags returns tags used by a logger. Tags are prepended to each line produced by a logger.
func (l *Logger) Tags() []string {
	r := make([]string, len(l.tags))
	copy(r, l.tags)
	return r
}

// Level returns log level used by a logger.
func (l *Logger) Level() Level {
	return l.level
}

// Output returns writer used by a logger.
func (l *Logger) Output() io.Writer {
	return l.output
}

// SetOutput changes Writer used to print logs. In contrast to other methods it modifies logger instance instead creating a new one.
func (l *Logger) SetOutput(output io.Writer) {
	l.output = output
}

// WithLevel creates new logger instance logging at specified level.
func (l *Logger) WithLevel(level Level) *Logger {
	n := *l
	n.level = level
	n.updateLinePrefix()
	return &n
}

// WithTags creates new logger instance with specified tags. Tags are prepended to each line produced by a logger.
func (l *Logger) WithTags(tags ...string) *Logger {
	n := *l
	n.tags = tags
	n.updateLinePrefix()
	return &n
}

// Printf writes log line. Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) *Logger {
	line := l.linePrefix + fmt.Sprint(v...) + "\033_klio_reset\033\\\n"
	l.output.Write([]byte(line))
	return l
}

// Printf writes log line. Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) *Logger {
	return l.Print(fmt.Sprintf(format, v...))
}

// Write prints input line by line.
func (l *Logger) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p)) // Scan lines
	for scanner.Scan() {
		l.Print(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(p), nil
}

// StandardLogger returns logger instance writing to the stdout. It writes using "info" level by default.
func StandardLogger() *Logger {
	return standardLogger
}

// ErrorLogger returns logger instance writing to the stderr. It writes using "error" level by default.
func ErrorLogger() *Logger {
	return errorLogger
}

// Spam writes a message at level Spam on the standard logger. Arguments are handled in the manner of fmt.Print.
func Spam(v ...interface{}) {
	standardLogger.WithLevel(SpamLevel).Print(v...)
}

// Debug writes a message at level Debug on the standard logger. Arguments are handled in the manner of fmt.Print.
func Debug(v ...interface{}) {
	standardLogger.WithLevel(DebugLevel).Print(v...)
}

// Verbose writes a message at level Verbose on the standard logger. Arguments are handled in the manner of fmt.Print.
func Verbose(v ...interface{}) {
	standardLogger.WithLevel(VerboseLevel).Print(v...)
}

// Info writes a message at level Info on the standard logger. Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {
	standardLogger.WithLevel(InfoLevel).Print(v...)
}

// Warn writes a message at level Warn on the standard logger. Arguments are handled in the manner of fmt.Print.
func Warn(v ...interface{}) {
	standardLogger.WithLevel(WarnLevel).Print(v...)
}

// Error writes a message at level Error on the standard logger. Arguments are handled in the manner of fmt.Print.
func Error(v ...interface{}) {
	standardLogger.WithLevel(ErrorLevel).Print(v...)
}

// Fatal writes a message at level Fatal on the standard logger. Arguments are handled in the manner of fmt.Print.
func Fatal(v ...interface{}) {
	standardLogger.WithLevel(FatalLevel).Print(v...)
}

// Spamf writes a message at level Spam on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Spamf(format string, v ...interface{}) {
	standardLogger.WithLevel(SpamLevel).Printf(format, v...)
}

// Debugf writes a message at level Debug on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	standardLogger.WithLevel(DebugLevel).Printf(format, v...)
}

// Verbosef writes a message at level Verbose on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Verbosef(format string, v ...interface{}) {
	standardLogger.WithLevel(VerboseLevel).Printf(format, v...)
}

// Infof writes a message at level Info on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	standardLogger.WithLevel(InfoLevel).Printf(format, v...)
}

// Warnf writes a message at level Warn on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	standardLogger.WithLevel(WarnLevel).Printf(format, v...)
}

// Errorf writes a message at level Error on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	standardLogger.WithLevel(ErrorLevel).Printf(format, v...)
}

// Fatalf writes a message at level Fatal on the standard logger. Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, v ...interface{}) {
	standardLogger.WithLevel(FatalLevel).Printf(format, v...)
}
