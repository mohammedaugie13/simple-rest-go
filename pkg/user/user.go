package user

import (
	"github.com/bitwyre/template-go/pkg/repository"
	"github.com/bitwyre/template-go/pkg/user/handler/rest"
	"github.com/bitwyre/template-go/pkg/user/router"
	"github.com/bitwyre/template-go/pkg/user/service"
	"github.com/gin-gonic/gin"
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
