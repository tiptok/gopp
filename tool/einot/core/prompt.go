package core

import (
	"context"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

type ChatTemplateConfig struct {
	FormatType schema.FormatType
	Templates  []schema.MessagesTemplate
}

// newChatTemplate component initialization function of node 'ChatTemplate1' in graph 'start'
func newChatTemplate(ctx context.Context) (ctp prompt.ChatTemplate, err error) {
	// TODO Modify component configuration here.
	config := &ChatTemplateConfig{
		FormatType: schema.FString,
		Templates:  make([]schema.MessagesTemplate, 0),
	}
	config.Templates = append(config.Templates, schema.MessagesPlaceholder("query", false))
	ctp = prompt.FromMessages(config.FormatType, config.Templates...)
	return ctp, nil
}

type ChatTemplate1Config struct {
	FormatType schema.FormatType
	Templates  []schema.MessagesTemplate
}

// newChatTemplate1 component initialization function of node 'ChatTemplate2' in graph 'start'
func newChatTemplate1(ctx context.Context) (ctp prompt.ChatTemplate, err error) {
	// TODO Modify component configuration here.
	config := &ChatTemplate1Config{
		FormatType: schema.FString,
		Templates:  make([]schema.MessagesTemplate, 0),
	}
	config.Templates = append(config.Templates, schema.MessagesPlaceholder("query", false))
	ctp = prompt.FromMessages(config.FormatType, config.Templates...)
	return ctp, nil
}
