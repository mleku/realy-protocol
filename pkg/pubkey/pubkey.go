package pubkey

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"io"
)

const Len = 43

type P struct{ ed25519.PublicKey }

func New(pk []byte) (p *P, err error) {
	if len(pk) != ed25519.PublicKeySize {
		err = errorf.E("invalid public key size: %d; require %d",
			len(pk), ed25519.PublicKeySize)
		return
	}
	p = &P{pk}
	return
}

func (p *P) Marshal(dst []byte) (result []byte, err error) {
	result = dst
	if p == nil || p.PublicKey == nil || len(p.PublicKey) == 0 {
		err = errorf.E("nil/zero length pubkey")
		return
	}
	if len(p.PublicKey) != ed25519.PublicKeySize {
		err = errorf.E("invalid public key length %d; require %d '%0x'",
			len(p.PublicKey), ed25519.PublicKeySize, p.PublicKey)
		return
	}
	buf := bytes.NewBuffer(result)
	w := base64.NewEncoder(base64.RawURLEncoding, buf)
	if _, err = w.Write(p.PublicKey); chk.E(err) {
		return
	}
	if err = w.Close(); chk.E(err) {
		return
	}
	result = append(buf.Bytes(), '\n')
	return
}

func (p *P) Unmarshal(data []byte) (rem []byte, err error) {
	rem = data
	if p == nil {
		err = errorf.E("can't unmarshal into nil types.T")
		return
	}
	if len(rem) < 2 {
		err = errorf.E("can't unmarshal nothing")
		return
	}
	for i := range rem {
		if rem[i] == '\n' {
			if i != Len {
				err = errorf.E("invalid encoded pubkey length %d; require %d '%0x'",
					i, Len, rem[:i])
				return
			}
			p.PublicKey = make([]byte, ed25519.PublicKeySize)
			if _, err = base64.RawURLEncoding.Decode(p.PublicKey, rem[:i]); chk.E(err) {
				return
			}
			rem = rem[i+1:]
			return
		}
	}
	err = io.EOF
	return
}
