package url

import (
	"testing"

	"protocol.realy.lol/pkg/separator"
)

func TestU_Marshal_Unmarshal(t *testing.T) {
	u := "https://example.com/path/to/resource"
	var err error
	var u1 *U
	if u1, err = New(u); chk.E(err) {
		t.Fatal(err)
	}
	var m1 []byte
	if m1, err = u1.Marshal(nil); chk.E(err) {
		t.Fatal(err)
	}
	m1 = separator.Add(m1)
	u2 := new(U)
	var rem []byte
	if rem, err = u2.Unmarshal(m1); chk.E(err) {
		t.Fatal(err)
	}
	if len(rem) > 0 {
		t.Fatalf("'%s' should be empty", string(rem))
	}
	if !u2.Equal(u1) {
		t.Fatalf("u1 should be equal to u2: '%s' != '%s'", u1, u2)
	}
}