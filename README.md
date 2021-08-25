# Klio logger for golang

This logger is meant to be used for building [klio](https://github.com/g2a-com/klio) commands (). It
writes logs decorated with
[control sequences interpreted by Klio](https://github.com/g2a-com/klio/blob/main/docs/output-handling.md).
It doesn't filter or modify messages besides that.

# Installation

```
go get github.com/g2a-com/klio-logger-go
```

# Quick Start

This logger comes with set of

```golang
package main

import (
	log "github.com/g2a-com/klio-logger-go"
)

func main() {
	log.Info("hello world")                  // Klio: [INFO] hello world
	log.Errorf("something went %s", "wrong") // Klio: [ERROR] something went wrong
}
```

For more complex cases, you may need to create custom logger. Since klio

### Custom

```golang
package main

import (
	log "github.com/g2a-com/klio-logger-go"
)

func main() {
	l := log.StandardLogger()
	l = l.WithTags("foo", "bar")      // Klio prepends tags to log lines
	l = l.WithLevel(log.VerboseLevel) // WithLevel and WithTags return new logger instead of modyfing the original one

	l.Print("hello world")            // Klio: [VERBOSE][FOO][BAR] hello world
	l.Printf("hello %s", "world")     // Klio: [VERBOSE][FOO][BAR] hello world
}
```
