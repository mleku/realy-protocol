package tags

import (
	"bytes"
	"fmt"

	"protocol.realy.lol/pkg/decimal"
	"protocol.realy.lol/pkg/separator"
	"protocol.realy.lol/pkg/tag"
)

const Sentinel = "tags:"

var SentinelBytes = []byte(Sentinel)

type tags []*tag.T

type T struct{ tags }

func New(v ...*tag.T) *T { return &T{tags: v} }

func (t *T) Marshal(dst []byte) (r []byte, err error) {
	r = dst
	r = append(r, Sentinel...)
	var l int
	if t != nil {
		l = len(t.tags)
	}
	if r, err = decimal.New(l).Marshal(r); chk.E(err) {
		return
	}
	r = separator.Add(r)
	if t != nil {
		for _, tt := range t.tags {
			if r, err = tt.Marshal(r); chk.E(err) {
				return
			}
		}
	}
	return
}

func (t *T) Unmarshal(data []byte) (rem []byte, err error) {
	if len(data) < len(Sentinel) {
		err = fmt.Errorf("bytes too short to contain tags")
		return
	}
	var d []byte
	if bytes.Equal(data[:len(Sentinel)], SentinelBytes) {
		d = data[len(Sentinel):]
	}
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
	for range l.N {
		tt := new(tag.T)
		if d, err = tt.Unmarshal(d); chk.E(err) {
			return
		}
		t.tags = append(t.tags, tt)
	}
	return
}