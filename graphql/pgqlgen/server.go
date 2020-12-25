package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tiptok/gopp/graphql/pgqlgen/resolvers"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tiptok/gopp/graphql/pgqlgen/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ginRun(port)
}

func httpRun(port string) {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/ui", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ginRun(port string) {
	svr := gin.New()
	svr.Use(gin.Recovery())

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))
	svr.Any("/ui", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	svr.Any("/query", gin.WrapH(h))
	svr.Run(":" + port)
}
