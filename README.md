# errors

[![CI](https://github.com/MJKWoolnough/errors/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/errors/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/errors.svg)](https://pkg.go.dev/vimagination.zapto.org/errors)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/errors)](https://goreportcard.com/report/vimagination.zapto.org/errors)

--
    import "vimagination.zapto.org/errors"

Package errors is a simple package with a few error related types.

## Highlights

 - Allows for const string error values.
 - Add stack traces to errors.
 - Add context messages to errors.

## Usage

```go
package main

import (
	"fmt"

	"vimagination.zapto.org/errors"
)

func main() {
	const ErrExample errors.Error = "my example error"

	err := errors.AddTrace(errors.WithContext("An error: ", ErrExample))

	fmt.Println(err)

	trace := err.(*errors.Trace)

	fmt.Println(trace.Traces[0].Function)

	// Output:
	// An error: my example error
	// main.main
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/errors
