package mime

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"strings"
	"testing"
)

func Test_multipart(t *testing.T) {
	var bufReader bytes.Buffer
	mpWriter := multipart.NewWriter(&bufReader)
	fw, err := mpWriter.CreateFormFile("file_a", "a.txt")
	if err != nil {
		t.Fatal(err)
		return
	}
	fw.Write([]byte("this is a txt"))
	mpWriter.WriteField("name", "tiprok")

	fwb, err := mpWriter.CreateFormFile("file_a", "b.txt")
	if err != nil {
		t.Fatal(err)
		return
	}
	fwb.Write([]byte("this is b txt"))
	mpWriter.Close()

	//t.Log("write ",mpWriter.FormDataContentType())
	t.Log("write", bufReader.String())

	mpReader := multipart.NewReader(&bufReader, mpWriter.Boundary())
	form, err := mpReader.ReadForm(100)
	if err != nil {
		t.Fatal(err)
	}
	for _, files := range form.File {
		for i := range files {
			file := files[i]
			if f, err := file.Open(); err != nil {
				return
			} else {
				bufTmp := bytes.NewBuffer(nil)
				bufTmp.ReadFrom(f)
				log.Println(file.Filename, " file:", bufTmp.String())
			}
		}
	}
	//log.Println("next part.")
	//for{
	//	if p,err :=mpReader.NextPart();(p!=nil && err==nil){
	//		log.Println(p.FileName())
	//	}else{
	//		break
	//	}
	//}
}

func TestMimeEncodedword(t *testing.T) {
	//邮件内容编码
	input := "我tip tok"
	out := mime.BEncoding.Encode("utf-8", input)
	outq := mime.QEncoding.Encode("utf-8", input)
	t.Log("base64:", base64.StdEncoding.EncodeToString([]byte(input)))
	t.Log(out)
	t.Log(outq)

	dec := new(mime.WordDecoder)
	if decout, err := dec.Decode(out); err == nil {
		t.Log("dec:", decout)
	}
	if decouthead, err := dec.DecodeHeader(out); err == nil {
		t.Log("dec head:", decouthead)
	}

	if decout, err := dec.Decode(outq); err == nil {
		t.Log("dec:", decout)
	}
	if decouthead, err := dec.DecodeHeader(outq); err == nil {
		t.Log("dec head:", decouthead)
	}
	array := [...][2]string{
		{"aa", "bb"},
		{"aa", "cc"},
	}
	for _, item := range array {
		t.Log(item[0], item[1])
	}
}

func TestFormatMediaType(t *testing.T) {
	type formatTest struct {
		typ    string
		params map[string]string
		want   string
	}
	var formatTests = []formatTest{
		{"noslash", map[string]string{"X": "Y"}, "noslash; x=Y"}, // e.g. Content-Disposition values (RFC 2183); issue 11289
		//{"foo bar/baz", nil, ""},
		{"foo/bar", map[string]string{"a": "av", "b": "bv", "c": "cv"}, "foo/bar; a=av; b=bv; c=cv"},
		{"attachment", map[string]string{"filename": "数据统计.png"}, "attachment; filename*=utf-8''%E6%95%B0%E6%8D%AE%E7%BB%9F%E8%AE%A1.png"},
	}
	for i, tt := range formatTests {
		got := mime.FormatMediaType(tt.typ, tt.params)
		if got != tt.want {
			t.Errorf("%d. FormatMediaType(%q, %v) = %q; want %q", i, tt.typ, tt.params, got, tt.want)
		}
		if got == "" {
			continue
		}
		typ, params, err := mime.ParseMediaType(got)
		if err != nil {
			t.Errorf("%d. ParseMediaType(%q) err: %v", i, got, err)
		}
		if typ != strings.ToLower(tt.typ) {
			t.Errorf("%d. ParseMediaType(%q) typ = %q; want %q", i, got, typ, tt.typ)
		}
		for k, v := range tt.params {
			k = strings.ToLower(k)
			if params[k] != v {
				t.Errorf("%d. ParseMediaType(%q) params[%s] = %q; want %q", i, got, k, params[k], v)
			}
		}
	}
}

func TestType(t *testing.T) {
	t.Log(mime.TypeByExtension(".jpeg"))
	t.Log(mime.TypeByExtension(".xml"))
	t.Log(mime.TypeByExtension("xml")) //not found
	t.Log(mime.TypeByExtension(".XmL"))
}

func TestQuotedprintable(t *testing.T) {
	for _, s := range []string{
		`=48=65=6C=6C=6F=2C=20=47=6F=70=68=65=72=73=21`,
		`invalid escape: <b style="font-size: 200%">hello</b>`,
		"Hello, Gophers! This symbol will be unescaped: =3D =3E =3F and this will be written in =\r\none line.",
	} {
		b, err := ioutil.ReadAll(quotedprintable.NewReader(strings.NewReader(s)))
		fmt.Printf("%s %v\n", b, err)
	}
}
