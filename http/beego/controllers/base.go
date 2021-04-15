package controllers

import (
	"github.com/beego/beego/v2/server/web/context"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/pkg/mybeego"
	"github.com/tiptok/gopp/pkg/protocol"
	"strings"
)

type Base struct {
	mybeego.ContextController
}

func (controller *Base) GetRequestHeader(ctx *context.Context) *protocol.RequestHeader {
	h := &protocol.RequestHeader{}

	if v := ctx.Input.GetData("x-mmm-id"); v != nil {
		h.UserId = int64(v.(int))
	}
	if v := ctx.Input.GetData("x-mmm-uname"); v != nil {
		h.UserName = v.(string)
	}
	h.Token = ctx.Input.Header("Authorization")
	if len(h.Token) > 0 && len(strings.Split(h.Token, " ")) > 1 {
		h.Token = strings.Split(h.Token, " ")[1]
	}
	h.BodyKeys = controller.BodyKeys(ctx, true)
	if v := ctx.Request.URL.Query(); len(v) > 0 {
		for k, _ := range v {
			h.BodyKeys = append(h.BodyKeys, common.CamelCase(k, true))
		}
	}
	return h
}
