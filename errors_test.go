package errors

import "testing"

func TestErrors(t *testing.T) {
	const (
		A Error = "A"
		B Error = "B"
	)
	var err error = A
	if err != A {
		t.Errorf("expecting err to match A")
	}
	if err == B {
		t.Errorf("not expecting err to match B")
	}
	err = B
	if err == A {
		t.Errorf("not expecting err to match A")
	}
	if err != B {
		t.Errorf("expecting err to match B")
	}
}
