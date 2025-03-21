package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino/schema"
	"github.com/tiptok/gopp/tool/einot/core"
)

func main() {
	einoTool, err := core.BuildStart(context.Background())
	if err != nil {
		panic(err)
	}

	out, err := einoTool.Invoke(context.Background(), map[string]any{"query": []*schema.Message{&schema.Message{
		Role:    schema.User,
		Content: "what is qwen?", //Content: "Beijing's weather this weekend",
	}}})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
