package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/graph/gqlerr"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
	"github.com/ezio1119/fishapp-api-gateway/pb"
	"google.golang.org/grpc/metadata"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	req := &pb.CreateUserReq{
		Data: &pb.CreateUserReq_Info{
			Info: &pb.CreateUserReqInfo{
				Email:        input.Email,
				Password:     input.Password,
				Name:         input.Name,
				Introduction: input.Introduction,
				Sex:          input.Sex,
			},
		},
	}

	stream, err := r.userClient.CreateUser(ctx)
	if err != nil {
		return nil, err
	}

	if err := stream.Send(req); err != nil {
		return nil, err
	}

	if input.Image != nil {
		for {
			buf := make([]byte, conf.C.Sv.ChunkDataSize)
			n, err := input.Image.File.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, gqlerr.InternalServerError("cannot read chunk to buffe: %s", err)
			}

			req = &pb.CreateUserReq{
				Data: &pb.CreateUserReq_ImageChunk{
					ImageChunk: buf[:n],
				},
			}

			if stream.Send(req); err != nil {
				return nil, fmt.Errorf("cannot send chunk to server: %w", err)
			}
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return &model.CreateUserPayload{
		User:      res.User,
		TokenPair: res.TokenPair,
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.UpdateUserPayload, error) {
	t, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError("missing token in 'Authorization' header: %s", err)
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"authorization": t}))

	req := &pb.UpdateUserReq{
		Data: &pb.UpdateUserReq_Info{
			Info: &pb.UpdateUserReqInfo{
				Email:        input.Email,
				Name:         input.Name,
				Introduction: input.Introduction,
			},
		},
	}

	stream, err := r.userClient.UpdateUser(ctx)
	if err != nil {
		return nil, err
	}

	if err := stream.Send(req); err != nil {
		return nil, err
	}

	if input.Image != nil {
		for {
			buf := make([]byte, conf.C.Sv.ChunkDataSize)
			n, err := input.Image.File.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, gqlerr.InternalServerError("cannot read chunk to buffe: %s", err)
			}

			req = &pb.UpdateUserReq{
				Data: &pb.UpdateUserReq_ImageChunk{
					ImageChunk: buf[:n],
				},
			}
		}

		if stream.Send(req); err != nil {
			return nil, fmt.Errorf("cannot send chunk to server: %w", err)
		}
	}

	u, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return &model.UpdateUserPayload{User: u}, nil
}

func (r *mutationResolver) UpdatePassword(ctx context.Context, input model.UpdatePasswordInput) (*model.UpdatePasswordPayload, error) {
	t, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError("missing token in 'Authorization' header: %s", err)
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"authorization": t}))

	if _, err := r.userClient.UpdatePassword(ctx, &pb.UpdatePasswordReq{
		OldPassword: input.OldPassword,
		NewPassword: input.NewPassword,
	}); err != nil {
		return nil, err
	}

	return &model.UpdatePasswordPayload{Success: true}, nil
}

func (r *mutationResolver) RefreshIDToken(ctx context.Context) (*model.RefreshIDTokenPayload, error) {
	t, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError("missing token in 'Authorization' header: %s", err)
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"authorization": t}))

	res, err := r.userClient.RefreshIDToken(ctx, &pb.RefreshIDTokenReq{})
	if err != nil {
		return nil, err
	}

	return &model.RefreshIDTokenPayload{TokenPair: res.TokenPair}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	res, err := r.userClient.Login(ctx, &pb.LoginReq{
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

	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"authorization": t}))

	if _, err := r.userClient.Logout(ctx, &pb.LogoutReq{}); err != nil {
		return nil, err
	}

	return &model.LogoutPayload{Success: true}, nil
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*pb.User, error) {
	t, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError("missing token in 'Authorization' header: %s", err)
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"authorization": t}))

	u, err := r.userClient.CurrentUser(ctx, &pb.CurrentUserReq{})
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *queryResolver) User(ctx context.Context, id int64) (*pb.User, error) {
	return r.userClient.GetUser(ctx, &pb.GetUserReq{Id: id})
}

func (r *userResolver) Posts(ctx context.Context, obj *pb.User) ([]*pb.Post, error) {
	res, err := r.postClient.ListPosts(ctx, &pb.ListPostsReq{Filter: &pb.ListPostsReq_Filter{
		UserId: obj.Id,
	}})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res.Posts, nil
}

func (r *userResolver) ApplyPosts(ctx context.Context, obj *pb.User) ([]*pb.ApplyPost, error) {
	res, err := r.postClient.ListApplyPosts(ctx, &pb.ListApplyPostsReq{
		Filter: &pb.ListApplyPostsReq_Filter{UserId: obj.Id},
	})
	if err != nil {
		return nil, err
	}

	return res.ApplyPosts, nil
}

func (r *userResolver) Image(ctx context.Context, obj *pb.User) (*pb.Image, error) {
	res, err := r.imageClient.ListImagesByOwnerID(ctx, &pb.ListImagesByOwnerIDReq{
		OwnerId:   obj.Id,
		OwnerType: pb.OwnerType_USER,
	})
	if err != nil {
		return nil, err
	}

	if len(res.Images) == 0 {
		return nil, err
	}

	return res.Images[0], nil
}

// User returns generated.UserResolver implementation.
func (r *resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *resolver }
