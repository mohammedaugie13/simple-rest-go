package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammedaugi13/simple-rest-go/pkg/repository"
	"github.com/mohammedaugi13/simple-rest-go/pkg/user"
)

type Module struct {
	UserModule *user.UserModule
}

func BootstrapModule(r *gin.Engine, repo *repository.Repository) *Module {
	return &Module{
		UserModule: user.NewUserModule(r, repo),
	}
}
