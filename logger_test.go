package logger_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	log "github.com/g2a-com/klio-logger-go"
)

func Example_basic() {
	log.Info("hello world")                  // Klio: [INFO] hello world
	log.Errorf("something went %s", "wrong") // Klio: [ERROR] something went wrong
}

func Example_tags() {
	l := log.StandardLogger().WithTags("foo", "bar")
	l.Print("tagged message") // Klio: [INFO][FOO][BAR] tagged message
}

func Example_stderr() {
	l := log.ErrorLogger().WithLevel(log.WarnLevel)
	l.Print("warning") // Klio: [WARN] warning
}

func Example_custom() {
	var b bytes.Buffer
	var l *log.Logger

	l = log.New(&b)                   // Logger can write to any io.Writer
	l = l.WithTags("foo", "bar")      // Klio prepends tags to log lines
	l = l.WithLevel(log.VerboseLevel) // WithLevel and WithTags return new logger instead of modyfing the original one

	l.Print("hello world")
	l.Printf("hello %s", "world")

	os.Stdout.Write(b.Bytes()) // Klio reads only from stdout and stderr, so after all you need to write your logs there
}

func TestParseLevel(t *testing.T) {
	var l log.Level
	var ok bool

	t.Run("parse valid levels", func(t *testing.T) {
		l, ok = log.ParseLevel("spam")
		assert.Equal(t, log.SpamLevel, l)
		assert.Equal(t, true, ok)

		l, ok = log.ParseLevel("debug")
		assert.Equal(t, log.DebugLevel, l)
		assert.Equal(t, true, ok)

		l, ok = log.ParseLevel("verbose")
		assert.Equal(t, log.VerboseLevel, l)
		assert.Equal(t, true, ok)

		l, ok = log.ParseLevel("info")
		assert.Equal(t, log.InfoLevel, l)
		assert.Equal(t, true, ok)

		l, ok = log.ParseLevel("warn")
		assert.Equal(t, log.WarnLevel, l)
		assert.Equal(t, true, ok)

		l, ok = log.ParseLevel("error")
		assert.Equal(t, log.ErrorLevel, l)
		assert.Equal(t, true, ok)

		l, ok = log.ParseLevel("fatal")
		assert.Equal(t, log.FatalLevel, l)
		assert.Equal(t, true, ok)
	})

	t.Run("parse levels with wrong letter case", func(t *testing.T) {
		l, ok = log.ParseLevel("SpAm")
		assert.Equal(t, log.SpamLevel, l)
		assert.Equal(t, true, ok)
	})

	t.Run("return DefaulLevel for invalid levels", func(t *testing.T) {
		l, ok = log.ParseLevel("unknown")
		assert.Equal(t, log.DefaultLevel, l)
		assert.Equal(t, false, ok)

		l, ok = log.ParseLevel("  spam")
		assert.Equal(t, log.DefaultLevel, l)
		assert.Equal(t, false, ok)
	})
}

func TestNew(t *testing.T) {
	var b bytes.Buffer
	l := log.New(&b)
	assert.IsType(t, &log.Logger{}, l)
}

func TestWithLevel(t *testing.T) {
	var b bytes.Buffer

	l1 := log.New(&b)
	l2 := l1.WithLevel(log.SpamLevel)
	l3 := l2.WithLevel(log.WarnLevel)

	assert.Equal(t, log.DefaultLevel, l1.Level())
	assert.Equal(t, log.SpamLevel, l2.Level())
	assert.Equal(t, log.WarnLevel, l3.Level())
}

func TestWithTags(t *testing.T) {
	var b bytes.Buffer

	l1 := log.New(&b)
	l2 := l1.WithTags("a", "b")
	l3 := l2.WithTags()

	l2.Tags()[0] = "xyz" // shouldn't affect logger tags

	assert.Equal(t, []string{}, l1.Tags())
	assert.Equal(t, []string{"a", "b"}, l2.Tags())
	assert.Equal(t, []string{}, l3.Tags())
}

func TestPrint(t *testing.T) {
	t.Run("print message with default level and no tags", func(t *testing.T) {
		var b bytes.Buffer
		log.New(&b).Print("foo")
		assert.Equal(t, "\033_klio_log_level \"info\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("print message with specified level and tags", func(t *testing.T) {
		var b bytes.Buffer
		log.New(&b).WithTags("a", "b", "c").WithLevel(log.SpamLevel).Print("foo")
		assert.Equal(t, "\033_klio_log_level \"spam\"\033\\\033_klio_tags [\"a\",\"b\",\"c\"]\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("ignore overwriten level and tags", func(t *testing.T) {
		var b bytes.Buffer
		log.New(&b).WithTags("a", "b", "c").WithLevel(log.SpamLevel).WithLevel(log.DefaultLevel).WithTags().Print("foo")
		assert.Equal(t, "\033_klio_log_level \"info\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("properly escape special characters in level and tags", func(t *testing.T) {
		var b bytes.Buffer
		log.New(&b).WithTags("\033\\").WithLevel(log.Level("\"")).Print("foo")
		assert.Equal(t, "\033_klio_log_level \"\\\"\"\033\\\033_klio_tags [\"\\u001b\\\\\"]\033\\foo\033_klio_reset\033\\\n", b.String())
	})
}

func TestPrintf(t *testing.T) {
	var b bytes.Buffer

	b.Reset()
	log.New(&b).Printf("%s", "foo")
	assert.Equal(t, "\033_klio_log_level \"info\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
}

func TestSetOutput(t *testing.T) {
	var b1 bytes.Buffer
	var b2 bytes.Buffer

	l := log.New(&b1)
	l.SetOutput(&b2)
	l.Print("foo")

	assert.Equal(t, "", b1.String())
	assert.Equal(t, "\033_klio_log_level \"info\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b2.String())
}

func TestStandardLogger(t *testing.T) {
	l := log.StandardLogger()

	assert.Equal(t, os.Stdout, l.Output())
	assert.Equal(t, log.Level("info"), l.Level())
	assert.Equal(t, []string{}, l.Tags())
}

func TestErrorLogger(t *testing.T) {
	l := log.ErrorLogger()

	assert.Equal(t, os.Stderr, l.Output())
	assert.Equal(t, log.Level("error"), l.Level())
	assert.Equal(t, []string{}, l.Tags())
}

func TestWriter(t *testing.T) {
	var b bytes.Buffer
	var w io.Writer = log.New(&b)

	n, err := w.Write([]byte("foo\nbar"))

	assert.NoError(t, err)
	assert.Equal(t, n, 7)
	assert.Equal(
		t,
		"\033_klio_log_level \"info\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n"+
			"\033_klio_log_level \"info\"\033\\\033_klio_tags []\033\\bar\033_klio_reset\033\\\n",
		b.String(),
	)
}

func TestConvenienceFunctions(t *testing.T) {
	var b bytes.Buffer

	log.StandardLogger().SetOutput(&b)
	defer log.StandardLogger().SetOutput(os.Stdout)

	t.Run("Spam", func(t *testing.T) {
		b.Reset()
		log.Spam("foo")
		assert.Equal(t, "\033_klio_log_level \"spam\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Debug", func(t *testing.T) {
		b.Reset()
		log.Debug("foo")
		assert.Equal(t, "\033_klio_log_level \"debug\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Verbose", func(t *testing.T) {
		b.Reset()
		log.Verbose("foo")
		assert.Equal(t, "\033_klio_log_level \"verbose\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Info", func(t *testing.T) {
		b.Reset()
		log.Info("foo")
		assert.Equal(t, "\033_klio_log_level \"info\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Warn", func(t *testing.T) {
		b.Reset()
		log.Warn("foo")
		assert.Equal(t, "\033_klio_log_level \"warn\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Error", func(t *testing.T) {
		b.Reset()
		log.Error("foo")
		assert.Equal(t, "\033_klio_log_level \"error\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Fatal", func(t *testing.T) {
		b.Reset()
		log.Fatal("foo")
		assert.Equal(t, "\033_klio_log_level \"fatal\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Spamf", func(t *testing.T) {
		b.Reset()
		log.Spamf("%s", "foo")
		assert.Equal(t, "\033_klio_log_level \"spam\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Debugf", func(t *testing.T) {
		b.Reset()
		log.Debugf("%s", "foo")
		assert.Equal(t, "\033_klio_log_level \"debug\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Verbosef", func(t *testing.T) {
		b.Reset()
		log.Verbosef("%s", "foo")
		assert.Equal(t, "\033_klio_log_level \"verbose\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Infof", func(t *testing.T) {
		b.Reset()
		log.Infof("%s", "foo")
		assert.Equal(t, "\033_klio_log_level \"info\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Warnf", func(t *testing.T) {
		b.Reset()
		log.Warnf("%s", "foo")
		assert.Equal(t, "\033_klio_log_level \"warn\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Errorf", func(t *testing.T) {
		b.Reset()
		log.Errorf("%s", "foo")
		assert.Equal(t, "\033_klio_log_level \"error\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})

	t.Run("Fatalf", func(t *testing.T) {
		b.Reset()
		log.Fatalf("%s", "foo")
		assert.Equal(t, "\033_klio_log_level \"fatal\"\033\\\033_klio_tags []\033\\foo\033_klio_reset\033\\\n", b.String())
	})
}
