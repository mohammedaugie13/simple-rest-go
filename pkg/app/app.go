package app

import (
	"github.com/bitwyre/template-go/pkg"
	"github.com/bitwyre/template-go/pkg/datastore/postgres"
	"github.com/bitwyre/template-go/pkg/handler/graph"
	"github.com/bitwyre/template-go/pkg/handler/rest"
	"github.com/bitwyre/template-go/pkg/repository"
	"github.com/bitwyre/template-go/pkg/service"
	"github.com/gin-gonic/gin"
)

func BootstrapApp(r *gin.Engine) {
	pg := postgres.NewClient()
	pg.AutoMigrate()

	repo := repository.NewRepository(pg.Db)
	rootService := service.NewService(repo)
	rootRestHandler := rest.NewRest(rootService)

	bootstrap := pkg.BootstrapModule(r, repo)
	gqlHandler := graph.NewGraphQL(rootService, bootstrap)
	NewRoutes(r, rootRestHandler, gqlHandler)
}
