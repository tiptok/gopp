package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tiptok/gocomm/pkg/mygin"
	"github.com/tiptok/gopp/pkg/protocol"
)

type BaseController struct {
	mygin.ContextController
}

func (controller *BaseController) GetRequestHeader(ctx *gin.Context) *protocol.RequestHeader {
	h := &protocol.RequestHeader{}
	return h
}
