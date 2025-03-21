package core

import (
	"context"

	"github.com/cloudwego/eino-ext/components/retriever/volc_vikingdb"
	"github.com/cloudwego/eino/components/retriever"
)

// newRetriever component initialization function of node 'Retriever' in graph 'start'
func newRetriever(ctx context.Context) (rtr retriever.Retriever, err error) {
	// TODO Modify component configuration here.
	config := &volc_vikingdb.RetrieverConfig{
		EmbeddingConfig: volc_vikingdb.EmbeddingConfig{}}
	embeddingIns11, err := newEmbedding(ctx)
	if err != nil {
		return nil, err
	}
	config.EmbeddingConfig.Embedding = embeddingIns11
	rtr, err = volc_vikingdb.NewRetriever(ctx, config)
	if err != nil {
		return nil, err
	}
	return rtr, nil
}
