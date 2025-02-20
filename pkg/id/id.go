package id

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"io"
)

const Len = 43

type T struct{ b []byte }

func New(id []byte) (p *T, err error) {
	if len(id) != ed25519.PublicKeySize {
		err = errorf.E("invalid public key size: %d; require %d",
			len(id), ed25519.PublicKeySize)
		return
	}
	p = &T{id}
	return
}

func (p *T) Marshal(d []byte) (r []byte, err error) {
	r = d
	if p == nil || p.b == nil || len(p.b) == 0 {
		err = errorf.E("nil/zero length pubkey")
		return
	}
	if len(p.b) != ed25519.PublicKeySize {
		err = errorf.E("invalid public key length %d; require %d '%0x'",
			len(p.b), ed25519.PublicKeySize, p.b)
		return
	}
	buf := bytes.NewBuffer(r)
	w := base64.NewEncoder(base64.RawURLEncoding, buf)
	if _, err = w.Write(p.b); chk.E(err) {
		return
	}
	if err = w.Close(); chk.E(err) {
		return
	}
	r = append(r, buf.Bytes()...)
	// r = append(buf.Bytes(), '\n')
	return
}

func (p *T) Unmarshal(data []byte) (r []byte, err error) {
	r = data
	if p == nil {
		err = errorf.E("can't unmarshal into nil types.T")
		return
	}
	if len(r) < 2 {
		err = errorf.E("can't unmarshal nothing")
		return
	}
	for i := range r {
		if r[i] == '\n' {
			if i != Len {
				err = errorf.E("invalid encoded pubkey length %d; require %d '%0x'",
					i, Len, r[:i])
				return
			}
			p.b = make([]byte, ed25519.PublicKeySize)
			if _, err = base64.RawURLEncoding.Decode(p.b, r[:i]); chk.E(err) {
				return
			}
			r = r[i+1:]
			return
		}
	}
	err = io.EOF
	return
}