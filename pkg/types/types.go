package types

import (
	"bytes"
	"io"
)

// A T is a type descriptor, that is terminated by a newline.
type T struct{ t []byte }

func New[V ~[]byte | ~string](t V) *T { return &T{[]byte(t)} }

func (t *T) Equal(t2 *T) bool { return bytes.Equal(t.t, t2.t) }

// Marshal append the T to a slice and appends a terminal newline, and returns
// the result.
func (t *T) Marshal(d []byte) (r []byte, err error) {
	if t == nil {
		return
	}
	r = append(d, t.t...)
	return
}

// Unmarshal expects an identifier followed by a newline. If the buffer ends
// without a newline an EOF is returned.
func (t *T) Unmarshal(d []byte) (r []byte, err error) {
	r = d
	if t == nil {
		err = errorf.E("can't unmarshal into nil types.T")
		return
	}
	if len(r) < 2 {
		err = errorf.E("can't unmarshal nothing")
		return
	}
	for i := range r {
		if r[i] == '\n' {
			// write read data up to the newline and return the remainder after
			// the newline.
			t.t = r[:i]
			r = r[i:]
			return
		}
	}
	// a T must end with a newline or an io.EOF is returned.
	err = io.EOF
	return
}