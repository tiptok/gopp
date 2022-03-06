package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestCookie(t *testing.T) {
	http.HandleFunc("/setCookie", setCookieHandler)
	http.HandleFunc("/getCookie", getCookieHandler)
	http.HandleFunc("/setMessage", setMessageHandler)
	http.HandleFunc("/getMessage", getMessageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "cookie_1",
		Value:    "ticctp",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "cookie_2",
		Value:    "tpccti",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Set("Set-Cookie",c2.String())
	w.Header().Add("Set-Cookie", c2.String())

	//或者
	//http.SetCookie(w,&c1)
	//http.SetCookie(w,&c2)
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)

	//或者
	//c1,err :=r.Cookie("cookie_1")
	//if err==nil{
	//	fmt.Fprintln(w,c1)
	//}
	//cs :=r.Cookies()
	//fmt.Fprintln(w,cs)
}

func setMessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := []byte("hello world")
	c := http.Cookie{
		Name:  "message",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func getMessageHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("message")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "message",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}
