package middleware

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/pb"

	"github.com/ezio1119/fishapp-api-gateway/graph/gqlerr"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
)

type graphQLMiddleware struct {
	chatClient pb.ChatServiceClient
	postClient pb.PostServiceClient
}

type GraphQLMiddleware interface {
	IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
	IsMember(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
	IsPostOwner(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
	IsApplyPostOwner(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
	IsUserOwner(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error)
}

func NewGraphQLMiddleware(cc pb.ChatServiceClient, pc pb.PostServiceClient) GraphQLMiddleware {
	return &graphQLMiddleware{cc, pc}
}

func (*graphQLMiddleware) IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fmt.Println("IsAuthenticated")
	t, err := getTokenFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError(err.Error())
	}

	c, err := getClaimsFromToken(t)
	if err != nil {
		return nil, gqlerr.AuthenticationError("token validation failed: %s", err)
	}

	if c.Subject != "id_token" {
		return nil, gqlerr.AuthenticationError("invalid tokentype: require id_token")
	}

	ctx = context.WithValue(ctx, model.JwtClaimsCtxKey, *c)

	return next(ctx)
}

func (m *graphQLMiddleware) IsMember(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fmt.Println("IsMember")
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	f := graphql.GetFieldContext(ctx)

	var isMember bool

	if pID, ok := f.Args["post_id"].(int64); ok {
		res, err := m.chatClient.IsMember(ctx, &pb.IsMemberReq{
			UserId:   uID,
			IsMember: &pb.IsMemberReq_PostId{PostId: pID},
		})
		if err != nil {
			return nil, err
		}

		isMember = res.Value
	}

	if in, ok := f.Args["input"].(model.CreateMessageInput); ok {
		res, err := m.chatClient.IsMember(ctx, &pb.IsMemberReq{
			UserId:   uID,
			IsMember: &pb.IsMemberReq_RoomId{RoomId: in.RoomID},
		})

		if err != nil {
			return nil, err
		}

		isMember = res.Value
	}

	if in, ok := f.Args["input"].(model.MessageAddedInput); ok {
		res, err := m.chatClient.IsMember(ctx, &pb.IsMemberReq{
			UserId:   uID,
			IsMember: &pb.IsMemberReq_RoomId{RoomId: in.RoomID},
		})

		if err != nil {
			return nil, err
		}

		isMember = res.Value
	}

	if !isMember {
		return nil, gqlerr.ForbiddenError("user_id=%d is not a member of the room")
	}

	return next(ctx)
}

func (m *graphQLMiddleware) IsPostOwner(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fmt.Println("IsPostOwner")
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	f := graphql.GetFieldContext(ctx)

	var pID int64
	switch f.Path().String() {
	case "updatePost":
		in, ok := f.Args["input"].(model.UpdatePostInput)
		if !ok {
			return nil, gqlerr.InternalServerError("field middleware IsPostOwner UpdatePostInput assertion failed")
		}
		pID = in.ID
	case "deletePost":
		in, ok := f.Args["input"].(model.DeletePostInput)
		if !ok {
			return nil, gqlerr.InternalServerError("field middleware IsPostOwner DeletePostInput assertion failed")
		}
		pID = in.ID
	}

	p, err := m.postClient.GetPost(ctx, &pb.GetPostReq{Id: pID})
	if err != nil {
		return nil, err
	}

	if p.UserId != uID {
		return nil, gqlerr.ForbiddenError("user_id=%d is not authorized because it is not the owner of post_id=%d", uID, pID)
	}

	return next(ctx)
}

func (m *graphQLMiddleware) IsApplyPostOwner(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	fmt.Println("IsApplyPostOwner")
	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	f := graphql.GetFieldContext(ctx)

	in, ok := f.Args["input"].(model.DeleteApplyPostInput)
	if !ok {
		return nil, gqlerr.InternalServerError("field middleware IsApplyPostOwner DeleteApplyPostInput assertion failed")
	}

	a, err := m.postClient.GetApplyPost(ctx, &pb.GetApplyPostReq{Id: in.ID})
	if err != nil {
		return nil, err
	}

	if a.UserId != uID {
		return nil, gqlerr.ForbiddenError("user_id=%d is not authorized because it is not the owner of apply_post_id=%d", uID, a.Id)
	}

	return next(ctx)
}

func (m *graphQLMiddleware) IsUserOwner(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	f := graphql.GetFieldContext(ctx)

	user, ok := f.Parent.Result.(*pb.User)
	if !ok {
		return nil, gqlerr.InternalServerError("field middleware IsUserOwner *pb.User assertion failed")
	}

	c, err := getClaimsFromCtx(ctx)
	if err != nil {
		return nil, gqlerr.AuthenticationError("authentication is required to see the %s with user_id=%d: %s", f.Field.Name, user.Id, err.Error())
	}

	uID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	if user.Id != uID {
		gqlerr.ForbiddenError("cant get email field because userid=%d is not the owner of user_id=%d", uID, user.Id)
	}

	return next(ctx)
}

func getClaimsFromToken(t string) (*model.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(t, &model.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if c, ok := token.Claims.(*model.JwtClaims); ok && token.Valid {
		return c, nil
	}

	return nil, err
}

func getClaimsFromCtx(ctx context.Context) (model.JwtClaims, error) {
	v := ctx.Value(model.JwtClaimsCtxKey)
	c, ok := v.(model.JwtClaims)
	if !ok {
		return model.JwtClaims{}, errors.New("failed to get jwt claims from context")
	}

	return c, nil
}

func getTokenFromCtx(ctx context.Context) (string, error) {
	v := ctx.Value(model.JwtTokenKey)

	token, ok := v.(string)
	if !ok {
		return "", errors.New("missing token in 'Authorization' header")
	}

	return token, nil
}

var publicKey *ecdsa.PublicKey

func init() {
	var err error
	publicKey, err = jwt.ParseECPublicKeyFromPEM([]byte(conf.C.Auth.PubJwtkey))
	if err != nil {
		panic(err)
	}
}
