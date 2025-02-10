// Package tag defines a format for event tags that follows the following rules:
//
// First field is the key, this is to be hashed using Blake2b and truncated to 8 bytes for indexing. These keys should
// not be long, and thus will not have any collisions as a truncated hash. The terminal byte of a key is the colon `:`
//
// Subsequent fields are separated by semicolon ';' and they can contain any data except a semicolon or newline.
//
// The tag is terminated by a newline.
package tag

import (
	"bytes"
)

type fields [][]byte

type T struct{ fields }

func New[V ~[]byte | ~string](v ...V) (t *T, err error) {
	t = new(T)
	var k []byte
	if k, err = ValidateKey([]byte(v[0])); err != nil {
		err = errorf.E("")
		return
	}
	v = v[1:]
	t.fields = append(t.fields, k)
	for i, val := range v {
		var b []byte
		if b, err = ValidateField(val, i); chk.E(err) {
			return
		}
		t.fields = append(t.fields, b)
	}
	return
}

// ValidateKey checks that the key is valid. Keys must be the same most language symbols:
//
// - first character is alphabetic [a-zA-Z]
// - subsequent characters can be alphanumeric and underscore [a-zA-Z0-9_]
//
// If the key is not valid this function returns a nil value.
func ValidateKey[V ~[]byte | ~string](key V) (k []byte, err error) {
	if len(key) < 1 {
		return
	}
	kb := []byte(key)
	switch {
	case kb[0] < 'a' && k[0] > 'z' || kb[0] < 'A' && kb[0] > 'Z':
		for i, b := range kb[1:] {
			switch {
			case (b > 'a' && b < 'z') || b > 'A' && b < 'Z' || b == '_' || b > '0' && b < '9':
			default:
				err = errorf.E("invalid character in tag key at index %d '%c': \"%s\"", i, b, kb)
				return
			}
		}
	}
	// if we got to here, the whole string is compliant
	k = kb
	return
}

func ValidateField[V ~[]byte | ~string](f V, i int) (k []byte, err error) {
	b := []byte(f)
	if bytes.Contains(b, []byte(";")) {
		err = errorf.E("key %d cannot contain ';': '%s'", i, b)
		return
	}
	if bytes.Contains(b, []byte("\n")) {
		err = errorf.E("key %d cannot contain '\\n': '%s'", i, b)
		return
	}
	// if we got to here, the whole string is compliant
	k = b
	return
}

func (t *T) Marshal(d []byte) (r []byte, err error) {
	r = d
	if len(t.fields) == 0 {
		return
	}
	for i, field := range t.fields {
		r = append(r, field...)
		if i == 0 {
			r = append(r, ':')
		} else if i == len(t.fields)-1 {
			r = append(r, '\n')
		} else {
			r = append(r, ';')
		}
	}
	return
}

func (t *T) Unmarshal(d []byte) (r []byte, err error) {
	var i int
	var v byte
	var dat []byte
	// first find the end
	for i, v = range d {
		if v == '\n' {
			dat, r = d[:i], d[i+1:]
			break
		}
	}
	if len(dat) == 0 {
		err = errorf.E("invalid empty tag")
		return
	}
	for i, v = range dat {
		if v == ':' {
			f := dat[:i]
			dat = dat[i+1:]
			t.fields = append(t.fields, f)
			break
		}
	}
	for len(dat) > 0 {
		for i, v = range dat {
			if v == ';' {
				t.fields = append(t.fields, dat[:i])
				dat = dat[i+1:]
				break
			}
			if i == len(dat)-1 {
				t.fields = append(t.fields, dat)
				return
			}
		}
	}
	return
}
