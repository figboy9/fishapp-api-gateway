package gqlerr

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorPresenter() graphql.ErrorPresenterFunc {
	return graphql.ErrorPresenterFunc(func(ctx context.Context, e error) *gqlerror.Error {
		s, ok := status.FromError(e) // grpc error
		if !ok {
			return graphql.DefaultErrorPresenter(ctx, e)
		}
		err := &gqlerror.Error{}
		switch s.Code() {
		case codes.InvalidArgument, codes.AlreadyExists, codes.NotFound:
			err = UserInputError(e.Error())
		case codes.Unauthenticated:
			err = AuthenticationError(e.Error())
		case codes.PermissionDenied:
			err = ForbiddenError(e.Error())
		default:
			err = InternalServerError(e.Error())
		}
		err.Extensions["grpc_code"] = s.Code().String()
		return err
	})
}
