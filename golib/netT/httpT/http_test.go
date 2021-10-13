package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

//FileTransport
func TestFileTransport(t *testing.T) {
	dname, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	fname := filepath.Join(dname, "foo.txt")
	defer os.Remove(dname)
	defer os.Remove(fname)
	err = ioutil.WriteFile(fname, []byte("Bar"), 0644)
	tr := &http.Transport{}
	tr.RegisterProtocol("file", http.NewFileTransport(http.Dir(dname)))
	fooURLs := []string{"file:///foo.txt", "file://../foo.txt"}
	c := http.Client{Transport: tr}
	for _, urlstr := range fooURLs {
		res, err := c.Get(urlstr)
		if res.StatusCode != http.StatusOK || err != nil {
			log.Fatal(res.StatusCode, err)
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		t.Log("read all:", string(body))
	}
}
