package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammedaugi13/simple-rest-go/pkg/user/handler/rest"
)

type Router struct {
	rest *rest.Rest
}

type MapRouter struct{}

func NewRouter(rest *rest.Rest) *Router {
	return &Router{
		rest,
	}
}

func (r *Router) RegisterRouter(gin *gin.Engine) {
	router := gin.Group("/user")
	router.GET("/health", r.rest.HealthCheck)
}
