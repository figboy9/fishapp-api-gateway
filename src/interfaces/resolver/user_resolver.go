package resolver

import (
	"context"
	"strconv"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
	"github.com/golang/protobuf/ptypes"
)

func (r *queryResolver) User(ctx context.Context, id string) (*graphql.User, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.UserInteractor.User(ctx, &user_grpc.ID{Id: intID})
}

func (r *mutationResolver) CreateUser(ctx context.Context, in gen.CreateUserInput) (*gen.UserWithToken, error) {
	req := &user_grpc.CreateReq{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	return r.UserInteractor.CreateUser(ctx, req)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, in gen.UpdateUserInput) (*graphql.User, error) {
	jwtClaims, err := getJwtCtx(ctx)
	if err != nil {
		return nil, err
	}
	req := &user_grpc.UpdateReq{
		Id:       jwtClaims.UserID,
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	return r.UserInteractor.UpdateUser(ctx, req)
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (bool, error) {
	jwtClaims, err := getJwtCtx(ctx)
	if err != nil {
		return false, err
	}
	return r.UserInteractor.DeleteUser(ctx, &user_grpc.ID{Id: jwtClaims.UserID})
}

func (r *mutationResolver) Login(ctx context.Context, in gen.LoginInput) (*gen.UserWithToken, error) {
	req := &user_grpc.LoginReq{
		Email:    in.Email,
		Password: in.Password,
	}
	return r.UserInteractor.Login(ctx, req)
}

func (r *postResolver) User(ctx context.Context, obj *graphql.Post) (*graphql.User, error) {
	intID, err := strconv.ParseInt(obj.UserID, 10, 64)
	if err != nil {
		return nil, err
	}
	return r.UserInteractor.User(ctx, &user_grpc.ID{Id: intID})
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	jwtClaims, err := getJwtCtx(ctx)
	if err != nil {
		return false, err
	}
	expiresAt := time.Unix(jwtClaims.ExpiresAt, 0)
	expiration := expiresAt.Sub(time.Now())
	expirationProto := ptypes.DurationProto(expiration)
	req := &user_grpc.AddBlackListReq{
		Jti:        jwtClaims.Jti,
		Expiration: expirationProto,
	}
	return r.UserInteractor.Logout(ctx, req)
}
func (r *mutationResolver) RefreshToken(ctx context.Context) (*graphql.TokenPair, error) {
	panic("not implemented")
}
