package netT

import (
	"encoding/json"
	"net/textproto"
	"sync"
	"testing"
)

func TestMimeHeader(t *testing.T) {
	var head textproto.MIMEHeader = textproto.MIMEHeader{}
	head.Add("aa-name", "tip")
	head.Add("cc-age", "tok")
	head.Add("aa-name", "2019")
	head.Add("test", "test")
	head.Add("!test", "test")
	head.Add("Upcase", "up")
	js, _ := json.Marshal(head)
	t.Log(string(js))
}

func TestPipeline(t *testing.T) {
	var pipe *textproto.Pipeline = new(textproto.Pipeline)
	var wg sync.WaitGroup
	event := pipe.Next()
	wg.Add(1)
	go func() {
		pipe.StartRequest(event)
		defer pipe.EndRequest(event)
		t.Log("on request")
		wg.Done()
	}()

	//go func(){
	//	pipe.StartResponse(event)
	//	defer pipe.EndResponse(event)
	//	t.Log("on response")
	//	wg.Done()
	//}()
	wg.Wait()
	t.Log("end.")
}
