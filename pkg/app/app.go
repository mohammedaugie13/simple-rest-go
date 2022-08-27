package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammedaugi13/simple-rest-go/pkg"
	"github.com/mohammedaugi13/simple-rest-go/pkg/datastore/postgres"
	"github.com/mohammedaugi13/simple-rest-go/pkg/handler/graph"
	"github.com/mohammedaugi13/simple-rest-go/pkg/handler/rest"
	"github.com/mohammedaugi13/simple-rest-go/pkg/repository"
	"github.com/mohammedaugi13/simple-rest-go/pkg/service"
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
