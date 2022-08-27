package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammedaugi13/simple-rest-go/pkg/handler/graph"
	"github.com/mohammedaugi13/simple-rest-go/pkg/handler/rest"
)

func NewRoutes(r *gin.Engine, rest *rest.Rest, gql *graph.GraphQL) {

	// GraphQL Routes
	gqlRoutes := r.Group("/graph")
	gqlRoutes.GET("/play", gql.GqlPlayground())
	gqlRoutes.POST("", gql.GqlServer())

	// Rest API Routes
	r.GET("/healthc", rest.HealthCheck)
}
