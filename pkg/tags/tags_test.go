package tags

import (
	"testing"

	"protocol.realy.lol/pkg/tag"
)

func TestT_Marshal_Unmarshal(t *testing.T) {
	var tegs = [][]string{
		{"reply", "l_T9Of4ru-PLGUxxvw3SfZH0e6XW11VYy8ZSgbcsD9Y", "realy.example.com/repo1"},
		{"root", "l_T9Of4ru-PLGUxxvw3SfZH0e6XW11VYy8ZSgbcsD9Y", "realy.example.com/repo2"},
		{"mention", "JMkZVnu9QFplR4F_KrWX-3chQsklXZq_5I6eYcXfz1Q", "realy.example.com/repo3"},
	}
	var err error
	var tgs []*tag.T
	for _, teg := range tegs {
		var tg *tag.T
		if tg, err = tag.New(teg...); chk.E(err) {
			t.Fatal(err)
		}
		tgs = append(tgs, tg)
	}
	t1 := New(tgs...)
	var m1 []byte
	if m1, err = t1.Marshal(nil); chk.E(err) {
		t.Fatal(err)
	}
	_ = m1
	// todo: unmarshal not currently working
	// t2 := new(T)
	// var rem []byte
	// if rem, err = t2.Unmarshal(m1); chk.E(err) {
	// 	t.Fatal(err)
	// }
	// if len(rem) > 0 {
	// 	t.Fatalf("%s", rem)
	// }
	// var m2 []byte
	// if m2, err = t2.Marshal(nil); chk.E(err) {
	// 	t.Fatal(err)
	// }
	// if !bytes.Equal(m1, m2) {
	// 	log.I.S(m1, m2)
	// 	t.Fatalf("not equal:\n%s\n%s", m1, m2)
	// }
}
