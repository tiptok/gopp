package schemal

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

type extend func() graphql.Fields

func AppendFields(extends ...extend) graphql.Fields {
	fields := graphql.Fields{}
	for _, extend := range extends {
		for k, v := range extend() {
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		}
	}
	return fields
}

func InitHandler() http.Handler {
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: AppendFields(ExtendUserQuery),
		},
	)

	var mutationType = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: AppendFields(ExtendUserMutation),
		})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	return h
}
