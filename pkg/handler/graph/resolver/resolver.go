package graph

import (
	"github.com/mohammedaugi13/simple-rest-go/pkg"
	"github.com/mohammedaugi13/simple-rest-go/pkg/handler/graph/generated"
	"github.com/mohammedaugi13/simple-rest-go/pkg/service"
)

type Resolver struct {
	module      *pkg.Module
	rootService *service.Service
}

func (r *Resolver) Mutation() generated.MutationResolver {
	//TODO implement me
	panic("implement me")
}

type MutationResolver struct{ *Resolver }
type QueryResolver struct{ *Resolver }

func NewResolver(rootService *service.Service, app *pkg.Module) *Resolver {
	return &Resolver{
		module:      app,
		rootService: rootService,
	}
}
