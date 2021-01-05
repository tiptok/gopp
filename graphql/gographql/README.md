## graphQL 库对比
![image](https://pic3.zhimg.com/80/v2-4a79ce03daf2fcdfcfc0607d1debe1fa_720w.jpg)

## graphQL-go 使用

### 1.声明schemal - object


声明对象 user


```
var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields:graphql.Fields{
		"id":&graphql.Field{
			Type: graphql.Int,
			Description: "user unique id",
		},
		"name":&graphql.Field{
			Type: graphql.String,                    // string 类型
			Description: "user name",
		},
		"roles":&graphql.Field{
			Type: graphql.NewList(Role),             // 数组类型
			Description: "user has role list",       // 属性描述
			Resolve: resolvers.UserResolvers.Roles,  // 定义解析方法
		},
		"status":&graphql.Field{
			Type: graphql.Int,
			Description: "user status (1:启用 2：禁用)",
		},
		"adminType":&graphql.Field{
			Type: graphql.Int,
			Description: "管理员类型 (1:超级管理员 2：普通用户)",
		},
		"createTime":&graphql.Field{
			Type: graphql.DateTime,
			Description: "创建时间",
		}
	},
})
```

### 2.定义schemal  query-mutation

定义fields

```
 graphql.Fields{
		"deleteUser":&graphql.Field{
			Type: User,             //类型
			Description: "delete user by id",  //描述
			Args: graphql.FieldConfigArgument{ //入参
				"id": &graphql.ArgumentConfig{ //声明参数 id
					Type: graphql.Int,
				},
			},
			Resolve: resolvers.UserResolvers.RemoveUser,//动作解析
		},
}
```

定义 query、mutation


```
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: AppendFields(ExtendUserQuery),
})
var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: AppendFields(ExtendUserMutation),
})
```

### 3.实例graphql服务


```
schema,_ :=graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
	Mutation: mutationType,
})

h:=handler.New(&handler.Config{
	Schema: &schema,
	Pretty: true,
	GraphiQL: false,
	Playground: true,
})

http.Handle("/query",schemal.InitHandler())
http.ListenAndServe(":8080"",nil)
```