package types

import (
	"io"
)

// A T is a type descriptor, that is terminated by a newline.
type T []byte

// Marshal append the T to a slice and appends a terminal newline, and returns
// the result.
func (t *T) Marshal(dst []byte) (result []byte, err error) {
	if t == nil {
		return
	}
	result = append(append(dst, []byte(*t)...), '\n')
	return
}

// Unmarshal expects an identifier followed by a newline. If the buffer ends
// without a newline an EOF is returned.
func (t *T) Unmarshal(data []byte) (rem []byte, err error) {
	rem = data
	if t == nil {
		err = errorf.E("can't unmarshal into nil types.T")
		return
	}
	if len(rem) < 2 {
		err = errorf.E("can't unmarshal nothing")
		return
	}
	for i := range rem {
		if rem[i] == '\n' {
			// write read data up to the newline and return the remainder after
			// the newline.
			*t = rem[:i]
			rem = rem[i+1:]
			return
		}
	}
	// a T must end with a newline or an io.EOF is returned.
	err = io.EOF
	return
}
