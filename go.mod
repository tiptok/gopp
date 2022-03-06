module github.com/tiptok/gopp

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/Laisky/laisky-blog-graphql v0.5.1
	github.com/beego/beego/v2 v2.0.1
	github.com/cloudwego/netpoll v0.0.3
	github.com/gin-gonic/gin v1.5.0
	github.com/go-pg/pg/v10 v10.7.7
	github.com/go-redis/redis/v8 v8.6.0
	github.com/golang/protobuf v1.4.3
	github.com/google/wire v0.5.0
	github.com/graphql-go/graphql v0.7.9
	github.com/graphql-go/handler v0.2.3
	github.com/klauspost/compress v1.11.4 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/stretchr/testify v1.7.0
	github.com/tal-tech/go-zero v1.1.1
	github.com/tiptok/gocomm v1.0.12
	github.com/vektah/gqlparser/v2 v2.1.0
	go.etcd.io/etcd v3.3.25+incompatible
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0
	gorm.io/driver/postgres v1.0.1
	gorm.io/gorm v1.20.1
)

//replace github.com/tiptok/gocomm v1.0.12 => /home/tiptok/go/src/gocomm
