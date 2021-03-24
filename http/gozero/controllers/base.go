package controllers

import (
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/pkg/mygozero"
	"github.com/tiptok/gopp/pkg/protocol"
	"net/http"
	"strings"
)

type Base struct {
	mygozero.RestBase
}

func (controller *Base) GetRequestHeader(r *http.Request) *protocol.RequestHeader {
	h := &protocol.RequestHeader{}

	if v := r.Context().Value("x-mmm-id"); v != nil {
		h.UserId = int64(v.(int))
	}
	if v := r.Context().Value("x-mmm-uname"); v != nil {
		h.UserName = v.(string)
	}
	h.Token = r.Header.Get("Authorization")
	if len(h.Token) > 0 && len(strings.Split(h.Token, " ")) > 1 {
		h.Token = strings.Split(h.Token, " ")[1]
	}
	h.BodyKeys = controller.BodyKeys(r, true)
	if v := r.URL.Query(); len(v) > 0 {
		for k, _ := range v {
			h.BodyKeys = append(h.BodyKeys, common.CamelCase(k, true))
		}
	}
	return h
}
