# paulog

This log module is a simple wrapper around the standard logging module. The logger prints pretty messages to the console.

## Installation

```bash
go get -u github.com/paust-team/paulog
```

## Usage

```go
package main

import (
    "github.com/paust-team/paulog"
)

var (
	logger = paulog.GetLogger("paulog.example.A")
)

func doSomething() error {
    return fmt.Errorf("some error")
}

func main() {
    logger.Infof("Hello World!")  // This will not be ignored
	err := doSomething()
	logger.Errorf("panic error: %v", err) // This will not be printed
	paulog.SetLevel("paulog.example", paulog.DEBUG)
    logger.Debugf("Print do this!")  // This will be printed
}

func init() {
	paulog.SetLevel("paulog", paulog.ERROR)
}
```

## License
MIT License