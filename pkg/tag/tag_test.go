package tag

import (
	"testing"
)

func TestT_Marshal_Unmarshal(t *testing.T) {
	var err error
	var t1 *T
	if t1, err = New("reply", "e:l_T9Of4ru-PLGUxxvw3SfZH0e6XW11VYy8ZSgbcsD9Y",
		"realy.example.com/repo"); chk.E(err) {
		t.Fatal(err)
	}
	var tb []byte
	if tb, err = t1.Marshal(nil); chk.E(err) {
		t.Fatal(err)
	}
	t2 := new(T)
	var rem []byte
	if rem, err = t2.Unmarshal(tb); chk.E(err) {
		t.Fatal(err)
	}
	if len(rem) > 0 {
		log.I.F("%s", rem)
		t.Fatal("remainder after tag should have been nothing")
	}
}
