package controller

import (
	"context"
	"encoding/json"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gopp/pkg/protocol"
)

type Base struct {
}

func (b *Base) JsonUnmarshal(dst interface{}, src interface{}) error {
	return common.UnmarshalFromString(common.JsonAssertString(src), dst)
}

func (controller *Base) GetRequestHeader(ctx context.Context, req interface{}) *protocol.RequestHeader {
	header := &protocol.RequestHeader{}

	header.BodyKeys = controller.BodyKeys(req)
	return header
}

func (controller *Base) BodyKeys(req interface{}) []string {
	var bodyKV map[string]json.RawMessage
	controller.JsonUnmarshal(req, &bodyKV)
	if len(bodyKV) == 0 {
		return []string{}
	}
	var list []string
	for k, _ := range bodyKV {
		list = append(list, common.CamelCase(k, true))
	}
	return list
}
