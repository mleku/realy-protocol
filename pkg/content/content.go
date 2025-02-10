package content

import (
	"bytes"
	"io"

	"protocol.realy.lol/pkg/decimal"
)

// C is raw content bytes of a message.
type C struct{ Content []byte }

// Marshal just writes the provided data with a `content:\n` prefix and adds a
// terminal newline.
func (c *C) Marshal(d []byte) (r []byte, err error) {
	r = append(d, "content:"...)
	if r, err = decimal.New(len(c.Content)).Marshal(r); chk.E(err) {
		return
	}
	r = append(r, '\n')
	// log.I.S(r)
	r = append(r, c.Content...)
	r = append(r, '\n')
	// log.I.S(r)
	return
}

var Prefix = "content:"

// Unmarshal expects the `content:<length>\n` prefix and stops at the second last
// newline. The data between the second last and last newline in the data is
// assumed to be a signature, but it could be anything in another use case.
//
// It is necessary that any non-content elements after the content must be
// parsed before returning to the content, because this is a
func (c *C) Unmarshal(d []byte) (r []byte, err error) {
	if !bytes.HasPrefix(d, []byte(Prefix)) {
		err = errorf.E("content prefix `content:' not found: '%s'", d[:len(Prefix)])
		return
	}
	// trim off the prefix.
	d = d[len(Prefix):]
	l := decimal.New(0)
	if d, err = l.Unmarshal(d); chk.E(err) {
		return
	}
	// and then there must be a newline
	if d[0] != '\n' {
		err = errorf.E("must be newline after content:<length>:\n%n", d)
		return
	}
	d = d[1:]
	// log.I.S(l.Uint64(), d)
	if len(d) < int(l.N) {
		err = io.EOF
		return
	}
	c.Content = d[:l.N]
	r = d[l.N:]
	if r[0] != '\n' {
		err = errorf.E("must be newline after content:<length>\\n, got '%s' %x", c.Content[len(c.Content)-1])
		return
	}
	r = r[1:]
	return
}
