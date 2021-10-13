package netT

import (
	"net/url"
	"testing"
)

func Test_parse(t *testing.T) {
	u, err := url.Parse("http://loacalhost:80/home/login?name=tip")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u.Scheme, u.Host, u.RawPath, u.RawQuery)

	urlStruct := url.URL{
		Scheme:   "https",
		Host:     "www.baidu.com",
		RawPath:  "/home/login",
		RawQuery: "name=tip",
	}
	q := urlStruct.Query()
	q.Set("year", "2019")
	urlStruct.RawQuery = q.Encode()
	t.Log(urlStruct.String())

	var value = url.Values{}
	value.Set("name", "tok")
	value.Set("age", "2018")
	urlStruct.RawQuery = value.Encode()
	t.Log(urlStruct.String())

	u1, _ := url.Parse("https://example.com/foo%2fbar")
	t.Log(u1.String())

	u2, _ := url.ParseRequestURI("https://example.com/foo=fbar")
	t.Log("ParseRequestURI: ", u2.String())
	t.Log("EscapedPath: ", u2.EscapedPath())
	t.Log("IsAbs: ", u2.IsAbs())

	u3, _ := url.Parse("../../..//search?q=dotnet")
	base, err := url.Parse("http://example.com/directory/")
	t.Log(base.ResolveReference(u3))
}
