package errors_test

import (
	"fmt"

	"vimagination.zapto.org/errors"
)

func Example() {
	const ErrExample errors.Error = "my example error"

	err := errors.AddTrace(errors.WithContext("An error: ", ErrExample))

	fmt.Println(err)

	trace := err.(*errors.Trace)

	fmt.Println(trace.Traces[0].Function)

	// Output:
	// An error: my example error
	// vimagination.zapto.org/errors_test.Example
}
