package tags

import (
	"protocol.realy.lol/pkg/tag"
)

type tags []*tag.T

type T struct{ tags }

func New(v ...*tag.T) *T { return &T{tags: v} }

func (t *T) Marshal(dst []byte) (r []byte, err error) {
	r = dst
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
	// todo: update for the lack of start/end markers
	// if len(data) < len(Sentinel) {
	// 	err = fmt.Errorf("bytes too short to contain tags")
	// 	return
	// }
	// var d []byte
	// if bytes.Equal(data[:len(Sentinel)], SentinelBytes) {
	// 	d = data[len(Sentinel):]
	// }
	// l := decimal.New(0)
	// if d, err = l.Unmarshal(d); chk.E(err) {
	// 	return
	// }
	// // and then there must be a newline
	// if d[0] != '\n' {
	// 	err = errorf.E("must be newline after content:<length>:\n%n", d)
	// 	return
	// }
	// d = d[1:]
	// for range l.N {
	// 	tt := new(tag.T)
	// 	if d, err = tt.Unmarshal(d); chk.E(err) {
	// 		return
	// 	}
	// 	t.tags = append(t.tags, tt)
	// }
	return
}
