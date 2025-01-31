package signature

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"io"
)

const Len = 86

type S struct{ Signature []byte }

func New(sig []byte) (p *S, err error) {
	if len(sig) != ed25519.SignatureSize {
		err = errorf.E("invalid signature size: %d; require %d",
			len(sig), ed25519.SignatureSize)
		return
	}
	p = &S{sig}
	return
}

func (p *S) Marshal(dst []byte) (result []byte, err error) {
	result = dst
	if p == nil || p.Signature == nil || len(p.Signature) == 0 {
		err = errorf.E("nil/zero length signature")
		return
	}
	if len(p.Signature) != ed25519.SignatureSize {
		err = errorf.E("invalid signature length %d; require %d '%0x'",
			len(p.Signature), ed25519.SignatureSize, p.Signature)
		return
	}
	buf := bytes.NewBuffer(result)
	w := base64.NewEncoder(base64.RawURLEncoding, buf)
	if _, err = w.Write(p.Signature); chk.E(err) {
		return
	}
	if err = w.Close(); chk.E(err) {
		return
	}
	result = append(buf.Bytes(), '\n')
	return
}

func (p *S) Unmarshal(data []byte) (rem []byte, err error) {
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
				err = errorf.E("invalid encoded signature length %d; require %d '%0x'",
					i, Len, rem[:i])
				return
			}
			p.Signature = make([]byte, ed25519.SignatureSize)
			if _, err = base64.RawURLEncoding.Decode(p.Signature, rem[:i]); chk.E(err) {
				return
			}
			rem = rem[i+1:]
			return
		}
	}
	err = io.EOF
	return
}
