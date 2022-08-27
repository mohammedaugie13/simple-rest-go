package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammedaugi13/simple-rest-go/pkg/repository"
	"github.com/mohammedaugi13/simple-rest-go/pkg/user/handler/rest"
	"github.com/mohammedaugi13/simple-rest-go/pkg/user/router"
	"github.com/mohammedaugi13/simple-rest-go/pkg/user/service"
)

type UserModule struct {
	Service     *service.Service
	RestHandler *rest.Rest
	Router      *router.Router
}

func NewUserModule(r *gin.Engine, repo *repository.Repository) *UserModule {
	services := service.NewService(repo)
	rests := rest.NewRest(services)
	routers := router.NewRouter(rests)
	routers.RegisterRouter(r)

	module := &UserModule{
		Service:     services,
		RestHandler: rests,
		Router:      routers,
	}

	return module
}
