package app

import (
	"github.com/bitwyre/template-go/pkg/handler/graph"
	"github.com/bitwyre/template-go/pkg/handler/rest"
	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.Engine, rest *rest.Rest, gql *graph.GraphQL) {

	// GraphQL Routes
	gqlRoutes := r.Group("/graph")
	gqlRoutes.GET("/play", gql.GqlPlayground())
	gqlRoutes.POST("", gql.GqlServer())

	// Rest API Routes
	r.GET("/healthc", rest.HealthCheck)
}
