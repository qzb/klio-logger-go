# Klio logger for golang

This logger is meant to be used for building [klio](https://github.com/g2a-com/klio) commands.
It writes logs decorated with control sequences [interpreted by Klio](https://github.com/g2a-com/klio/blob/main/docs/output-handling.md).

```go
import (
  log "github.com/g2a-com/klio-logger-go"
)

log.Info("hello world")                                                     // [INFO] hello world
log.StandardLogger().WithLevel("spam").WithTags("foo", "bar").Print("tags") // [SPAM][FOO][BAR] tags
```

## Functions

### func [Debug](/logger.go#L166)

`func Debug(v ...interface{})`

Debug writes a message at level Debug on the standard logger. Arguments are handled in the manner of fmt.Print.

### func [Debugf](/logger.go#L201)

`func Debugf(format string, v ...interface{})`

Debugf writes a message at level Debug on the standard logger. Arguments are handled in the manner of fmt.Printf.

### func [Error](/logger.go#L186)

`func Error(v ...interface{})`

Error writes a message at level Error on the standard logger. Arguments are handled in the manner of fmt.Print.

### func [Errorf](/logger.go#L221)

`func Errorf(format string, v ...interface{})`

Errorf writes a message at level Error on the standard logger. Arguments are handled in the manner of fmt.Printf.

### func [Fatal](/logger.go#L191)

`func Fatal(v ...interface{})`

Fatal writes a message at level Fatal on the standard logger. Arguments are handled in the manner of fmt.Print.

### func [Fatalf](/logger.go#L226)

`func Fatalf(format string, v ...interface{})`

Fatalf writes a message at level Fatal on the standard logger. Arguments are handled in the manner of fmt.Printf.

### func [Info](/logger.go#L176)

`func Info(v ...interface{})`

Info writes a message at level Info on the standard logger. Arguments are handled in the manner of fmt.Print.

### func [Infof](/logger.go#L211)

`func Infof(format string, v ...interface{})`

Infof writes a message at level Info on the standard logger. Arguments are handled in the manner of fmt.Printf.

### func [Spam](/logger.go#L161)

`func Spam(v ...interface{})`

Spam writes a message at level Spam on the standard logger. Arguments are handled in the manner of fmt.Print.

### func [Spamf](/logger.go#L196)

`func Spamf(format string, v ...interface{})`

Spamf writes a message at level Spam on the standard logger. Arguments are handled in the manner of fmt.Printf.

### func [Verbose](/logger.go#L171)

`func Verbose(v ...interface{})`

Verbose writes a message at level Verbose on the standard logger. Arguments are handled in the manner of fmt.Print.

### func [Verbosef](/logger.go#L206)

`func Verbosef(format string, v ...interface{})`

Verbosef writes a message at level Verbose on the standard logger. Arguments are handled in the manner of fmt.Printf.

### func [Warn](/logger.go#L181)

`func Warn(v ...interface{})`

Warn writes a message at level Warn on the standard logger. Arguments are handled in the manner of fmt.Print.

### func [Warnf](/logger.go#L216)

`func Warnf(format string, v ...interface{})`

Warnf writes a message at level Warn on the standard logger. Arguments are handled in the manner of fmt.Printf.

## Types

### type [Level](/logger.go#L21)

`type Level string`

Level type.

#### Constants

```golang
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
```

#### func [ParseLevel](/logger.go#L57)

`func ParseLevel(s string) (level Level, ok bool)`

ParseLevel converts level name to Level. It is case insensitive, returns DefaultLevel if value cannot be converted.

### type [Logger](/logger.go#L66)

`type Logger struct { ... }`

Logger.

#### func [ErrorLogger](/logger.go#L156)

`func ErrorLogger() *Logger`

ErrorLogger returns logger instance writing to the stderr.

#### func [New](/logger.go#L74)

`func New(output io.Writer) *Logger`

New creates new instance of Logger.

#### func [StandardLogger](/logger.go#L151)

`func StandardLogger() *Logger`

StandardLogger returns logger instance writing to the stdout.

#### func (*Logger) [Level](/logger.go#L108)

`func (l *Logger) Level() Level`

Level returns log level used by a logger.

#### func (*Logger) [Output](/logger.go#L113)

`func (l *Logger) Output() io.Writer`

Output returns writer used by a logger.

#### func (*Logger) [Print](/logger.go#L139)

`func (l *Logger) Print(v ...interface{ ... }) *Logger`

Printf writes log line. Arguments are handled in the manner of fmt.Print.

#### func (*Logger) [Printf](/logger.go#L146)

`func (l *Logger) Printf(format string, v ...interface{ ... }) *Logger`

Printf writes log line. Arguments are handled in the manner of fmt.Printf.

#### func (*Logger) [SetOutput](/logger.go#L118)

`func (l *Logger) SetOutput(output io.Writer)`

SetOutput changes Writer used to print logs. In contrast to other methods it modifies logger instance instead creating a new one.

#### func (*Logger) [Tags](/logger.go#L101)

`func (l *Logger) Tags() []string`

Tags returns tags used by a logger. ags are prepended to each line produced by a logger.

#### func (*Logger) [WithLevel](/logger.go#L123)

`func (l *Logger) WithLevel(level Level) *Logger`

WithLevel creates new logger instance logging at specified level.

#### func (*Logger) [WithTags](/logger.go#L131)

`func (l *Logger) WithTags(tags ...string) *Logger`

WithTags creates new logger instance with specified tags. Tags are prepended to each line produced by a logger.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
