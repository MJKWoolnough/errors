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
	if err.Error() != "A" {
		t.Errorf("expecting error string \"A\" got %q", err.Error())
	}
	err = B
	if err == A {
		t.Errorf("not expecting err to match A")
	}
	if err != B {
		t.Errorf("expecting err to match B")
	}
	if err.Error() != "B" {
		t.Errorf("expecting error string \"B\" got %q", err.Error())
	}
	if c := WithContext("CONTEXT", A); Underlying(c) != A {
		t.Errorf("expecting error %q, got %q", A, c)
	}
}

func TestItobs(t *testing.T) {
	tests := []struct {
		Input  int
		Output []byte
	}{
		{0, []byte{'0'}},
		{1, []byte{'1'}},
		{2, []byte{'2'}},
		{10, []byte{'1', '0'}},
		{11, []byte{'1', '1'}},
		{32, []byte{'3', '2'}},
		{100, []byte{'1', '0', '0'}},
		{-1, []byte{'-', '1'}},
		{-2, []byte{'-', '2'}},
		{-10, []byte{'-', '1', '0'}},
		{-100, []byte{'-', '1', '0', '0'}},
	}
	for n, test := range tests {
		o := itobs(test.Input)
		if string(o) != string(test.Output) {
			t.Errorf("test %d: expecting %s, got %s", n+1, test.Output, o)
		}
	}
}
