package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"testing"
)

func TestDump(t *testing.T) {
	//1.http.Request
	Req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Host:   "www.google.com",
			Path:   "/search",
		},
		ProtoMajor:       1,
		ProtoMinor:       1,
		TransferEncoding: []string{"chunked"},
	}
	Req.Body = ioutil.NopCloser(bytes.NewReader([]byte("abdldoelddd")))
	dumpWithBody, _ := httputil.DumpRequest(Req, true)
	t.Log(string(dumpWithBody))

	dumpWithOutBody, _ := httputil.DumpRequest(Req, false)
	t.Log(string(dumpWithOutBody))

	//2.http data
	reqData := "POST /v2/api/?login HTTP/1.1\r\n" +
		"Host: passport.myhost.com\r\n" +
		"Content-Length:5\r\n" +
		"\r\nkey1=name1&key2=name2"
	req2, _ := http.ReadRequest(bufio.NewReader(strings.NewReader(reqData)))
	dumpWithBody_req2, _ := httputil.DumpRequest(req2, true)
	t.Log(string(dumpWithBody_req2))

	dumpWithOutBody_req2, _ := httputil.DumpRequest(req2, false)
	t.Log(string(dumpWithOutBody_req2))

	//3.response
	res := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: 3,
		Body:          ioutil.NopCloser(strings.NewReader("foo")),
	}
	rspDump, _ := httputil.DumpResponse(res, true)
	t.Log(string(rspDump))
}
