package schemal

import (
	"github.com/graphql-go/graphql"
	"github.com/tiptok/gopp/graphql/gographql/resolvers"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "user unique id",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "user name",
		},
		"phone": &graphql.Field{
			Type:        graphql.String,
			Description: "user phone num",
		},
		"roles": &graphql.Field{
			Type:        graphql.NewList(Role),
			Description: "user has role list",
			Resolve:     resolvers.UserResolvers.Roles,
		},
		"status": &graphql.Field{
			Type:        graphql.Int,
			Description: "user status (1:启用 2：禁用)",
		},
		"adminType": &graphql.Field{
			Type:        graphql.Int,
			Description: "管理员类型 (1:超级管理员 2：普通用户)",
		},
		"createTime": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "创建时间",
		},
		"updateTime": &graphql.Field{
			Type:        graphql.DateTime,
			Description: "更新时间",
		},
	},
})

var UserList = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserList",
	Fields: graphql.Fields{
		"total": &graphql.Field{
			Type:        graphql.Int,
			Description: "user unique id",
		},
		"users": &graphql.Field{
			Type:        graphql.NewList(User),
			Description: "user unique id",
		},
	},
})

func ExtendUserQuery() graphql.Fields {
	return graphql.Fields{
		"user": &graphql.Field{
			Type:        User,
			Description: "Get user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolvers.UserResolvers.User,
		},
		"users": &graphql.Field{
			Type:        UserList,
			Description: "Get user list",
			Args: graphql.FieldConfigArgument{
				"limit": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"offset": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"searchByText": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"sortById": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolvers.UserResolvers.Users,
		},
	}
}

func ExtendUserMutation() graphql.Fields {
	return graphql.Fields{
		"deleteUser": &graphql.Field{
			Type:        User,
			Description: "delete user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolvers.UserResolvers.RemoveUser,
		},
		"createUser": &graphql.Field{
			Type:        User,
			Description: "update user by id",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "user name",
				},
				"phone": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "user phone num",
				},
				"passwd": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "user account password",
				},
				"roles": &graphql.ArgumentConfig{
					Type:        graphql.NewList(graphql.Int),
					Description: "user has role list",
				},
				"status": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "user status (1:启用 2：禁用)",
				},
				"adminType": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "管理员类型 (1:超级管理员 2：普通用户)",
				},
			},
			Resolve: resolvers.UserResolvers.CreateUser,
		},
		"updateUser": &graphql.Field{
			Type:        User,
			Description: "update user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"name": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "user name",
				},
				"phone": &graphql.ArgumentConfig{
					Type:        graphql.String,
					Description: "user phone num",
				},
				"roles": &graphql.ArgumentConfig{
					Type:        graphql.NewList(graphql.Int),
					Description: "user has role list",
				},
				"status": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "user status (1:启用 2：禁用)",
				},
				"adminType": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "管理员类型 (1:超级管理员 2：普通用户)",
				},
			},
			Resolve: resolvers.UserResolvers.UpdateUser,
		},
	}
}
