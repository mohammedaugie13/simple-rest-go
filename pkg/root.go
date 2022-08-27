package pkg

import (
	"github.com/bitwyre/template-go/pkg/repository"
	"github.com/bitwyre/template-go/pkg/user"
	"github.com/gin-gonic/gin"
)

type Module struct {
	UserModule *user.UserModule
}

func BootstrapModule(r *gin.Engine, repo *repository.Repository) *Module {
	return &Module{
		UserModule: user.NewUserModule(r, repo),
	}
}
