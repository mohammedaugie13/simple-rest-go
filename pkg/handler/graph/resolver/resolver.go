package graph

import (
	"github.com/bitwyre/template-go/pkg"
	"github.com/bitwyre/template-go/pkg/handler/graph/generated"
	"github.com/bitwyre/template-go/pkg/service"
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
