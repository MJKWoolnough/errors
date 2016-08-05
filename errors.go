package errors

import "runtime"

type Error string

func (e Error) Error() string {
	return string(e)
}

type Call struct {
	File, Function string
	LineNum        int
}

func (c Call) String() string {
	return c.File + ": (" + string(itobs(c.LineNum)) + ")" + c.Function
}

type Trace struct {
	error
	Traces []Call
}

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

type Underlyer interface {
	Underlying() error
}

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
