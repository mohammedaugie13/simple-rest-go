package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bitwyre/template-go/pkg/handler/graph/generated"
	"github.com/bitwyre/template-go/pkg/handler/graph/model"
	"github.com/bitwyre/template-go/pkg/helper"
)

func (r *QueryResolver) InfoQuery(ctx context.Context, input *model.InfoQueryInput) (*model.InfoQuery, error) {
	result, err := r.module.UserModule.Service.UserAuthService.CheckUserStatus(input)
	if err != nil {
		helper.GqlError(ctx, &helper.GraphQLError{
			Message: err,
			ErrCode: "USER_NOT_FOUND",
		})

		return nil, nil
	}

	return result, nil
}

func (r *Resolver) Query() generated.QueryResolver { return &QueryResolver{r} }
