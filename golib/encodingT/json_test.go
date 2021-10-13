package encodingT

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

/*json.RawMessage 使用*/
type jsonIn struct {
	Head string
	/*
	   1.首字母需要大写才能编码
	   2.用于结构体重的一个字段的格式未知，等确定类型后进行二次Unmarshal
	*/
	Body *json.RawMessage
}
type bodyMessage struct {
	Message string
	Sender  string
	To      string
}

type jsonOut struct {
	Head string
	Body interface{}
}

func TestRawMessage(t *testing.T) {
	jOut := &jsonOut{
		Head: "bodyMessage",
		Body: bodyMessage{Message: "201802 Send msg", Sender: "tip", To: "tok"},
	}
	byteOut, err := json.Marshal(jOut)
	if err != nil {
		t.Fatal(err.Error())
	}
	jIn := jsonIn{}
	log.Printf("%s", string(byteOut))
	err = json.Unmarshal(byteOut, &jIn)
	if err != nil {
		t.Fatal(err.Error())
	}
	log.Println("head:", jIn.Head)
	if jIn.Body != nil {
		log.Printf("UnMarshal bytes:%v", jIn.Body)
	}

	/*动态反射实例*/
	var message bodyMessage
	tType := reflect.TypeOf(message)
	var iMessage interface{}
	iMessage = reflect.New(tType).Interface()

	/*二次解析*/
	err = json.Unmarshal(*jIn.Body, &iMessage)
	if err != nil {
		t.Fatal(err.Error())
	}
	log.Println("二次解析:", iMessage)
}

//reflect.New(t).Interface()
