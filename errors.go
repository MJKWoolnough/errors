// Package errors is a simple package with a few error related types
package errors

import "runtime"

// Error represents a constant error string
type Error string

// Error returns the error string
func (e Error) Error() string {
	return string(e)
}

// Call represents where a particular error was created
type Call struct {
	File, Function string
	LineNum        int
}

// String returns a human-friendly representation of the call site
func (c Call) String() string {
	return c.File + ": (" + string(itobs(c.LineNum)) + ")" + c.Function
}

// Trace represents the call stack for an error
type Trace struct {
	error
	Traces []Call
}

// AddTrace wraps an error with a call stack
func AddTrace(e error) error {
	var trace [100]uintptr
	num := runtime.Callers(2, trace[:])
	traces := make([]Call, num)
	for i := 0; i < num; i++ {
		f := runtime.FuncForPC(trace[i])
		file, ln := f.FileLine(trace[i])
		traces[i] = Call{
			File:     file,
			Function: f.Name(),
			LineNum:  ln,
		}
	}
	return &Trace{
		error:  e,
		Traces: traces,
	}
}

// Trace returns a byte slice containing a description of the call stack
func (t Trace) Trace() []byte {
	var buf []byte
	for _, c := range t.Traces {
		buf = append(buf, c.File...)
		buf = append(buf, ':', ' ', '(')
		buf = append(buf, itobs(c.LineNum)...)
		buf = append(buf, ')', ' ')
		buf = append(buf, c.Function...)
		buf = append(buf, '\n')
	}
	return buf
}

// Underlyer is used to get the underlying error for a wrapped error
type Underlyer interface {
	Underlying() error
}

// Underlying returns the underlying error
func (t Trace) Underlying() error {
	if u, ok := t.error.(Underlyer); ok {
		return u.Underlying()
	}
	return t.error
}

func itobs(i int) []byte {
	if i == 0 {
		return []byte{'0'}
	}
	var neg = false
	if i < 0 {
		neg = true
		i = -i
	}
	pos := 21
	var b [21]byte
	for ; i > 0; i /= 10 {
		pos--
		b[pos] = byte(i%10) + '0'
	}
	if neg {
		pos--
		b[pos] = '-'
	}
	return b[pos:]
}

type contextual struct {
	context string
	error
}

func WithContext(context string, err error) error {
	if err == nil {
		return nil
	}
	return &contextual{
		context: context,
		error:   err,
	}
}

func (c *contextual) Error() string {
	return c.context + c.error.Error()
}

func (c *contextual) Underlying() error {
	return c.error
}
