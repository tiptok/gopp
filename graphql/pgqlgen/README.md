## 一、快速开始

### 初始化
```
需要提前安装 gqlgen(https://github.com/99designs/gqlgen) 到go/bin目录
install gqlgen

$ mkdir gqlgen-todos
$ cd gqlgen-todos
$ go mod init github.com/[username]/gqlgen-todos
$ gqlgen init
$ gplgen generate
```

[官方文档-快速开始](https://gqlgen.com/getting-started/)


### 自定义字段rosolvers

延后查询 roles字段
```
user(input:{id:1}){
    phone,
    name,
    status,
    roles{
      id,
      roleName
    }
  }
```

假设我们有一个user对象，里面的roles对象是通过user里面的字段，延后查询获取到，
那么我们需要新建一个 user.go模型文件,覆盖model_gen.go里面生成的user对象

```
type Users struct {
	Roles      []*int  `json:"roles"`
}

// 角色模型
type Role struct {
	ID         *int      `json:"id"`
	RoleName   *string   `json:"roleName"`
}
```
定义schemal文件

```
type Users{
  roles: [Role!]
}
```
执行指令 gqlgen generate , 会生成一个usersResolever对象，我们只需要实现Roles这个方法，就可以实现延后查询的功能

```
type usersResolver struct{ *Resolver }

func (r *usersResolver) Roles(ctx context.Context, obj *model.Users) ([]*model.Role, error) {

    // 在此处实现roles获取的逻辑
}
```

方法二 通过配置文件 增加 resolver: true

```
models
  ...
  Users:
    fields:
      roles:
        resolver: true
```

方法三 定义schemal字段为resolver

```
type Users @goModel(model: "github.com/tiptok/gopp/pkg/domain.Users") {
    id: ID!
    name: String
    phone: String
    roles: [Role!] @goField(forceResolver: true)
}
```




### 分割 schemal.resovers.go 到多个文件

系统初始化的schemal.graphqls 文件，我们只需要把新的模型建立到新的graphqls文件中去，例如：
新建/user/users.graphqls 根目录执行 ==gqlgen generate==,会生成对应的
==users.resovers.go== 文件，以后的逻辑实现写在users.resovers.go文件中即可

需要修改配置文件 gqlgen.yml文件，加入新的schemal声明文件路径
```
schema:
  - graph/*.graphqls
  - graph/users/*.graphqls
```


### shememal.graphqls文件太大

分离schemal文件到自定义 custome.graphqls文件中去
通过extend关键字，扩展Query 和 Mutation，就不需要把所有的方法都放一个graphql文件里面声明

```
extend type Query {
  menu(input: menuAccessInput): [Access!]
}
extend type Mutation {
  createMenu(input: menuAccessInput): [Access!]
}
```


### 复用已经声明的模型

方法一

引入domain 里面的 ClientVersion  ClientPackageInfo
修改gqlgen.yml

```
models:
  Date:
    model: github.com/tiptok/godevp/pkg/port/graphql/graph/libs.Datetime
  ClientVersion:
    model: github.com/tiptok/godevp/pkg/domain.ClientVersion
  ClientPackageInfo:
    model: github.com/tiptok/godevp/pkg/domain.ClientPackageInfo
```
修改xxx.graphql,声明类型 ClientVersion ClientPackageInfo ，（映射到graphql类型的时候会忽略字段类型大小写）

```
## 声明Date 时间类型
scalar Date


type ClientVersion{
  id: Int
  commiter: String
  projectName: String
  version: String
  title: String
  remark: String
  clientPackageInfo: [ClientPackageInfo!]
  createTime: Date
}

type ClientPackageInfo{
  FileName: String
  Path: String
}

type Query {
  clientVersion(id :Int):ClientVersion!
}
```

实现clientVersion获取的方法

```
func (r *queryResolver) ClientVersion(ctx context.Context, id *int) (*domain1.ClientVersion, error) {
	svr:=clientVersionService.NewClientVersionService(nil)
	m,err:=svr.GetClientVersion(&query2.GetClientVersionQuery{Id: int64(*id)})
	if err!=nil || m==nil{
		return nil, err
	}
	return m.(*domain1.ClientVersion),nil
}
```
自定义的Date类型需要实现接口

UnmarshalGQL(vi interface{}) (err error)

MarshalGQL(w io.Writer)
```
type Datetime struct {
	t time.Time
}

const TimeLayout = "2006-01-02T15:04:05.000Z"
func (d *Datetime) UnmarshalGQL(vi interface{}) (err error) {
	v, ok := vi.(string)
	if !ok {
		return fmt.Errorf("unknown type of Datetime: `%+v`", vi)
	}
	if d.t, err = time.Parse(TimeLayout, v); err != nil {
		return err
	}

	return nil
}

func (d Datetime) MarshalGQL(w io.Writer) {
	if _, err := w.Write(appendQuote([]byte(d.t.Format(TimeLayout)))); err != nil {
	}
}
```

方法二

使用标签 @goModel
```
type Todo @goModel(model: "github.com/NateScarlet/issue-repro/model.Todo") {
  id: ID!
  text: String!
}
```



## 二、查询 query

```
query{
  user(input:{id:1}){
    phone,
    name,
    status,
    roles{
      id,
      roleName
    }
  }
  users(input:{limit:10,offset:0}){
      phone,
      name
  }
}
```

## 三、变化 mutation

```
mutation createTodo {
  createTodo(input:{text:"todo", userId:"1"}) {
    user {
      id
    }
    text
    done
  }
}
```

## 四、坑

### 1、Int
gqlgen 在序列化 int64 的时候精度有问题，候建议转为字符串来传递。


