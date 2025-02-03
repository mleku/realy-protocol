package url

import (
	"bytes"
	"net/url"
)

type U struct{ uu []byte }

// New creates a new URL codec.C from the provided URL, and validates it.
func New[V ~string | []byte](ur V) (uu *U, err error) {
	uu = new(U)
	var UU *url.URL
	if UU, err = url.Parse(string(ur)); chk.E(err) {
		return
	} else {
		// if it is valid, store it
		uu.uu = []byte(UU.String())
	}
	return
}

func (u *U) String() string   { return string(u.uu) }
func (u *U) Bytes() []byte    { return u.uu }
func (u *U) Equal(u2 *U) bool { return bytes.Equal(u.uu, u2.uu) }

// Marshal a URL, use New to ensure it is valid beforehand. Appends a terminal
// newline.
func (u *U) Marshal(dst []byte) (result []byte, err error) {
	result = append(append(dst, u.uu...), '\n')
	return
}

// Unmarshal decodes a URL and validates it is a proper URL.
func (u *U) Unmarshal(data []byte) (rem []byte, err error) {
	rem = data
	for i, v := range rem {
		if v == '\n' {
			u.uu = rem[:i]
			rem = rem[i+1:]
			break
		}
	}
	// validate the URL and return error if not valid.
	if _, err = url.Parse(string(u.uu)); chk.E(err) {
		return
	}
	return
}
