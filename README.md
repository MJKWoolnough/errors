# errors
--
    import "vimagination.zapto.org/errors"

Package errors is a simple package with a few error related types.

## Usage

#### func  AddTrace

```go
func AddTrace(e error) error
```
AddTrace wraps an error with a call stack.

#### func  New

```go
func New(str string) error
```
New returns an error that returns the given string.

#### func  Unwrap

```go
func Unwrap(err error) error
```
Unwrap repeatedly called checks for an underlying error to returned the original
wrapped error.

#### func  WithContext

```go
func WithContext(context string, err error) error
```
WithContext wraps an error, adding textural context to the error message.

The wrapped error can be accessed via the Unwrap method.

A nil error will not be wrapped.

#### type Call

```go
type Call struct {
	File, Function string
	LineNum        int
}
```

Call represents where a particular error was created.

#### func (Call) String

```go
func (c Call) String() string
```
String returns a human-friendly representation of the call site.

#### type Error

```go
type Error string
```

Error represents a constant error string.

#### func (Error) Error

```go
func (e Error) Error() string
```
Error returns the error string.

#### type Trace

```go
type Trace struct {
	Traces []Call
}
```

Trace represents the call stack for an error.

#### func (Trace) Trace

```go
func (t Trace) Trace() []byte
```
Trace returns a byte slice containing a description of the call stack.

#### func (Trace) Unwrap

```go
func (t Trace) Unwrap() error
```
Unwrap returns the underlying error.

#### type Wrapper

```go
type Wrapper interface {
	Unwrap() error
}
```

Wrapper is used to get the underlying error for a wrapped error.
