package content

import (
	"bytes"
)

// C is raw content bytes of a message. This can contain anything but when it is
// unmarshalled it is assumed that the last line (content between the second
// last and last line break) is not part of the content, as this is where the
// signature is placed.
//
// The only guaranteed property of an encoded content.C is that it has two
// newline characters, one at the very end, and a second one before it that
// demarcates the end of the actual content. It can be entirely binary and mess
// up a terminal to render the unsanitized possible control characters.
type C struct{ Content []byte }

// Marshal just writes the provided data with a `content:\n` prefix and adds a
// terminal newline.
func (c *C) Marshal(dst []byte) (result []byte, err error) {
	result = append(append(append(dst, "content:\n"...), c.Content...), '\n')
	return
}

var Prefix = "content:\n"

// Unmarshal expects the `content:\n` prefix and stops at the second last
// newline. The data between the second last and last newline in the data is
// assumed to be a signature but it could be anything in another use case.
func (c *C) Unmarshal(data []byte) (rem []byte, err error) {
	if !bytes.HasPrefix(data, []byte("content:\n")) {
		err = errorf.E("content prefix `content:\\n' not found: '%s'", data[:len(Prefix)+1])
		return
	}
	// trim off the prefix.
	data = data[len(Prefix):]
	// check that there is a last newline.
	if data[len(data)-1] != '\n' {
		err = errorf.E("input data does not end with newline")
		return
	}
	// we start at the second last, previous to the terminal newline byte.
	lastPos := len(data) - 2
	for ; lastPos >= len(Prefix); lastPos-- {
		// the content ends at the byte before the second last newline byte.
		if data[lastPos] == '\n' {
			break
		}
	}
	c.Content = data[:lastPos]
	// return the remainder after the content-terminal newline byte.
	rem = data[lastPos+1:]
	return
}
