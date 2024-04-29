package fuzz

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

func Reverse(s string) (string, error) {

	b := []rune(s)
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}

func TestFmt(t *testing.T) {
	type SessionUser struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var su = SessionUser{
		Name: "su",
		Age:  40,
	}
	gob.Register(SessionUser{})
	data := fmt.Sprintf("%v", su)
	var toSu = SessionUser{}
	json.Unmarshal([]byte(data), &toSu)
	assert.Equal(t, toSu.Name, su.Name)
}
