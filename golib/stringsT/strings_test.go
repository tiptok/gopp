package stringsT

import (
	"strings"
	"testing"
	"time"
)

func TestStrings(t *testing.T) {
	var sBuild strings.Builder
	t.Log(sBuild.Cap())
	sBuild.Grow(4)
	sBuild.WriteString("tiptok")
	t.Log(sBuild.String(), sBuild.Len())
}

func TestEqual(t *testing.T) {
	input := struct {
		A string
		B string
		C string
	}{
		A: "ab我",
		B: "ac",
		C: "ab我",
	}
	if strings.EqualFold(input.A, input.B) != false {
		t.Fatal("equal:", input.A, input.B)
	}
	if !strings.EqualFold(input.A, input.C) {
		t.Fatal("equal", input.A, input.C)
	}
	if strings.Compare(input.A, input.C) != 0 {
		t.Fatal("equal", input.A, input.C)
	}
}

func TestStringify(t *testing.T) {
	var nilPointer *string

	var tests = []struct {
		in  interface{}
		out string
	}{
		// basic types
		{"foo", `"foo"`},
		{123, `123`},
		{1.5, `1.5`},
		{false, `false`},
		{
			[]string{"a", "b"},
			`["a" "b"]`,
		},
		{
			struct {
				A []string
			}{nil},
			// nil slice is skipped
			`{}`,
		},
		{
			struct {
				A string
			}{"foo"},
			// structs not of a named type get no prefix
			`{A:"foo"}`,
		},

		// pointers
		{nilPointer, `<nil>`},
		{String("foo"), `"foo"`},
		{Int(123), `123`},
		{Bool(false), `false`},
		{
			[]*string{String("a"), String("b")},
			`["a" "b"]`,
		},

		// actual GitHub structs
		{
			Timestamp{time.Date(2006, time.January, 02, 15, 04, 05, 0, time.UTC)},
			`stringsT.Timestamp{2006-01-02 15:04:05 +0000 UTC}`,
		},
		{
			&Timestamp{time.Date(2006, time.January, 02, 15, 04, 05, 0, time.UTC)},
			`stringsT.Timestamp{2006-01-02 15:04:05 +0000 UTC}`,
		},
	}

	for i, tt := range tests {
		s := Stringify(tt.in)
		if s != tt.out {
			t.Errorf("%d. Stringify(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
	t.Log(Stringify(struct {
		A int
		B string
		C []int
		D map[string]int
		E Timestamp
	}{1, "2", []int{3, 3, 3}, map[string]int{"4": 4}, Timestamp{time.Now()}}))
}
