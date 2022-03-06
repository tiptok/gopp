package encodingT

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestBase64(t *testing.T) {
	input := []byte("http://127.0.0.1:8080/tiptok")
	base64.StdEncoding.EncodedLen(10)
	t.Log(base64.StdEncoding.EncodeToString(input))
	var buf bytes.Buffer
	e := base64.NewEncoder(base64.URLEncoding, &buf)
	e.Write(input)
	e.Close()
	t.Log(buf.String())
}

func TestUTF8(t *testing.T) {
	input := "abc你我他"
	t.Log(len(input))
	for i := 0; i < len(input); i++ {
		r, n := utf8.DecodeRuneInString(input[i:])
		t.Log(fmt.Sprintf("%c %d", r, n))
	}
}
