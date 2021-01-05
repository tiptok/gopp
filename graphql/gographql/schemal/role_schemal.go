package schemal

import "github.com/graphql-go/graphql"

var Role = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"roleName": &graphql.Field{
			Type: graphql.String,
		},
	},
})
