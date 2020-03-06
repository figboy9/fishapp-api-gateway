package errors

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/status"
)

func ErrorPresenter() graphql.ErrorPresenterFunc {
	return graphql.ErrorPresenterFunc(func(ctx context.Context, e error) *gqlerror.Error {
		s, ok := status.FromError(e) // grpc error
		if ok {
			return &gqlerror.Error{
				Message:    e.Error(),
				Path:       graphql.GetFieldContext(ctx).Path(),
				Extensions: map[string]interface{}{"code": s.Code().String(), "type": "grpc"},
			}
		}
		return graphql.DefaultErrorPresenter(ctx, e)
	})
}
