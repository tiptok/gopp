package core

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
)

// newChatModel component initialization function of node 'ChatModel1' in graph 'start'
func newChatModel(ctx context.Context) (cm model.ChatModel, err error) {
	// TODO Modify component configuration here.
	key := os.Getenv("DASHSCOPE_API_KEY")
	config := &openai.ChatModelConfig{
		Model:   "qwen-plus",
		APIKey:  key,
		BaseURL: "https://dashscope.aliyuncs.com/compatible-mode/v1/",
	}
	cm, err = openai.NewChatModel(ctx, config)
	if err != nil {
		return nil, err
	}
	return cm, nil
}
