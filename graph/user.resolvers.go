package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/graph/gqlerr"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/grpc/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
	"google.golang.org/grpc/metadata"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	// _, err := r.pubsubClient.CreateTopic(ctx, "sample_topic")
	// if err != nil {
	// 	fmt.Println("エラーだおん")
	// 	return nil, err
	// }

	t := r.pubsubClient.Topic("sample_topic")
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Message ああああああ"),
	})
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Println("エラーだ")
		return nil, err
	}
	fmt.Println("published id=", id)
	sub := r.pubsubClient.Subscription("sample_topic")
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
		cancel()
	})
	if err != nil {
		return nil, fmt.Errorf("Receive: %w", err)
	}
	fmt.Println("終了")
	return nil, fmt.Errorf("おワーリ")

	// res, err := r.authClient.CreateUser(ctx, &auth_grpc.CreateUserReq{
	// 	Email:    input.Email,
	// 	Password: input.Password,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// return &model.CreateUserPayload{
	// 	User:      res.User,
	// 	TokenPair: res.TokenPair,
	// }, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.UpdateUserPayload, error) {
	t, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError("missing token in 'Authorization' header: %s", err)
	}
	u, err := r.authClient.UpdateUser(
		metadata.AppendToOutgoingContext(ctx, "authorization", t),
		&auth_grpc.UpdateUserReq{
			Email:       input.Email,
			OldPassword: input.OldPassword,
			Password:    input.Password,
		})
	if err != nil {
		return nil, err
	}
	return &model.UpdateUserPayload{User: u}, nil
}

func (r *mutationResolver) RefreshIDToken(ctx context.Context) (*model.RefreshIDTokenPayload, error) {
	t, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError("missing token in 'Authorization' header: %s", err)
	}
	res, err := r.authClient.RefreshIDToken(
		metadata.AppendToOutgoingContext(ctx, "authorization", t),
		&auth_grpc.RefreshIDTokenReq{},
	)
	if err != nil {
		return nil, err
	}
	return &model.RefreshIDTokenPayload{TokenPair: res.TokenPair}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	res, err := r.authClient.Login(ctx, &auth_grpc.LoginReq{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	return &model.LoginPayload{User: res.User, TokenPair: res.TokenPair}, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (*model.LogoutPayload, error) {
	t, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError("missing token in 'Authorization' header: %s", err)
	}
	if _, err := r.authClient.Logout(
		metadata.AppendToOutgoingContext(ctx, "authorization", t),
		&auth_grpc.LogoutReq{},
	); err != nil {
		return nil, err
	}
	return &model.LogoutPayload{Success: true}, nil
}

func (r *queryResolver) User(ctx context.Context) (*auth_grpc.User, error) {
	t, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError("missing token in 'Authorization' header: %s", err)
	}
	u, err := r.authClient.GetUser(
		metadata.AppendToOutgoingContext(ctx, "authorization", t),
		&auth_grpc.GetUserReq{},
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userResolver) Posts(ctx context.Context, obj *auth_grpc.User) ([]*post_grpc.Post, error) {
	res, err := r.postClient.ListPosts(ctx, &post_grpc.ListPostsReq{Filter: &post_grpc.ListPostsReq_Filter{
		UserId: obj.Id,
	}})
	if err != nil {
		return nil, err
	}
	return res.Posts, nil
}

func (r *userResolver) ApplyPosts(ctx context.Context, obj *auth_grpc.User) ([]*post_grpc.ApplyPost, error) {
	res, err := r.postClient.ListApplyPosts(ctx, &post_grpc.ListApplyPostsReq{
		Filter: &post_grpc.ListApplyPostsReq_Filter{UserId: obj.Id},
	})
	if err != nil {
		return nil, err
	}
	return res.ApplyPosts, nil
}

func (r *userResolver) Profile(ctx context.Context, obj *auth_grpc.User) (*profile_grpc.Profile, error) {
	p, err := r.profileClient.GetProfile(ctx, &profile_grpc.GetProfileReq{UserId: obj.Id})
	if err != nil {
		return nil, err
	}
	return p, nil
}

// User returns generated.UserResolver implementation.
func (r *resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *resolver }
