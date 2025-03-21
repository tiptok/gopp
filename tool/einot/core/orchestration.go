package core

import (
	"context"

	"github.com/cloudwego/eino/compose"
)

func BuildStart(ctx context.Context) (r compose.Runnable[any, any], err error) {
	const (
		Retriever     = "Retriever"
		ChatTemplate1 = "ChatTemplate1"
		ChatTemplate2 = "ChatTemplate2"
		ChatModel1    = "ChatModel1"
	)
	g := compose.NewGraph[any, any]()
	//retrieverKeyOfRetriever, err := newRetriever(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//_ = g.AddRetrieverNode(Retriever, retrieverKeyOfRetriever)
	chatTemplate1KeyOfChatTemplate, err := newChatTemplate(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddChatTemplateNode(ChatTemplate1, chatTemplate1KeyOfChatTemplate)
	//chatTemplate2KeyOfChatTemplate, err := newChatTemplate1(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//_ = g.AddChatTemplateNode(ChatTemplate2, chatTemplate2KeyOfChatTemplate)
	chatModel1KeyOfChatModel, err := newChatModel(ctx)
	if err != nil {
		return nil, err
	}
	_ = g.AddChatModelNode(ChatModel1, chatModel1KeyOfChatModel)
	//_ = g.AddEdge(compose.START, Retriever)
	_ = g.AddEdge(compose.START, ChatTemplate1)
	//_ = g.AddEdge(compose.START, ChatTemplate2)
	_ = g.AddEdge(ChatModel1, compose.END)
	//_ = g.AddEdge(Retriever, ChatTemplate1)
	//_ = g.AddEdge(Retriever, ChatTemplate2)
	_ = g.AddEdge(ChatTemplate1, ChatModel1)
	//_ = g.AddEdge(ChatTemplate2, ChatModel1)
	r, err = g.Compile(ctx, compose.WithGraphName("start"), compose.WithNodeTriggerMode(compose.AnyPredecessor))
	if err != nil {
		return nil, err
	}
	return r, err
}
