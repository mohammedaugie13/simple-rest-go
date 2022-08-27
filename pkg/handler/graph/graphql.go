package graph

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bitwyre/template-go/pkg"
	"github.com/bitwyre/template-go/pkg/handler/graph/generated"
	graph "github.com/bitwyre/template-go/pkg/handler/graph/resolver"
	"github.com/bitwyre/template-go/pkg/lib"
	"github.com/bitwyre/template-go/pkg/service"
	"github.com/gin-gonic/gin"
)

type GraphQL struct {
	module      *pkg.Module
	rootService *service.Service
}

var graphqlServer = "/graph"
var graphqlPlayground = "/graph/play"

func NewGraphQL(rootService *service.Service, module *pkg.Module) *GraphQL {
	log.Println("🟢 GraphQL Server: ", graphqlServer)
	if lib.AppConfig.App.Env != "prod" {
		log.Println("🟢 GraphQL Playground: ", graphqlPlayground)
	} else {
		log.Println("🔴 GraphQL Playground: Disabled")
	}

	return &GraphQL{
		module:      module,
		rootService: rootService,
	}
}

func (g *GraphQL) GqlServer() gin.HandlerFunc {
	resolver := graph.NewResolver(g.rootService, g.module)
	c := generated.Config{Resolvers: resolver}
	c.Directives.Binding = Binding
	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func (g *GraphQL) GqlPlayground() gin.HandlerFunc {
	h := playground.Handler("GraphQL", graphqlServer)
	return func(ctx *gin.Context) {
		if lib.AppConfig.App.Env != "prod" {
			h.ServeHTTP(ctx.Writer, ctx.Request)
			return
		}
	}
}
