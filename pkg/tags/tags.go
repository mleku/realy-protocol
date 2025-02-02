package tags

import (
	"bytes"
	"fmt"

	"protocol.realy.lol/pkg/tag"
)

const Sentinel = "tags:\n"

var SentinelBytes = []byte(Sentinel)

type tags []*tag.T

type T struct{ tags }

func New(v ...*tag.T) *T { return &T{tags: v} }

func (t *T) Marshal(dst []byte) (result []byte, err error) {
	result = dst
	result = append(result, Sentinel...)
	for _, tt := range t.tags {
		if result, err = tt.Marshal(result); chk.E(err) {
			return
		}
	}
	result = append(result, '\n')
	return
}

func (t *T) Unmarshal(data []byte) (rem []byte, err error) {
	if len(data) < len(Sentinel) {
		err = fmt.Errorf("bytes too short to contain tags")
		return
	}
	var dat []byte
	if bytes.Equal(data[:len(Sentinel)], SentinelBytes) {
		dat = data[len(Sentinel):]
	}
	if len(dat) < 1 {
		return
	}
	for len(dat) > 0 {
		if len(dat) == 1 && dat[0] == '\n' {
			break
		}
		// log.I.S(dat)
		tt := new(tag.T)
		if dat, err = tt.Unmarshal(dat); chk.E(err) {
			return
		}
		t.tags = append(t.tags, tt)
	}
	return
}
