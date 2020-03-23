package middleware

import (
	"context"
	"crypto/ecdsa"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/ezio1119/fishapp-api-gateway/graph/model"
)

func (*middleware) AuthMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver, authAPI bool) (res interface{}, err error) {
	if authAPI {
		return next(ctx)
	}
	t, err := getTokenCtx(ctx)
	if err != nil {
		return nil, err
	}

	c, err := getClaimsToken(t)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, model.JwtClaimsCtxKey, *c)

	return next(ctx)
}

// func setJwtClaimsCtx(ctx context.Context, c *model.JwtClaims) (context.Context, error) {
// 	userID, err := strconv.ParseInt(c.User.ID, 10, 64)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return context.WithValue(ctx, model.JwtClaimsCtxKey, graph.JwtClaims{
// 		UserID:    userID,
// 		Jti:       c.Id,
// 		ExpiresAt: c.ExpiresAt,
// 	}), nil
// }

func getClaimsToken(t string) (*model.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(t, &model.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if c, ok := token.Claims.(*model.JwtClaims); ok && token.Valid {
		return c, nil
	}
	return nil, err
}

func getTokenCtx(ctx context.Context) (string, error) {
	v := ctx.Value(model.JwtTokenKey)
	token, ok := v.(string)
	if !ok {
		return "", &gqlerror.Error{
			Message: "missing token in 'Authorization' header",
			Extensions: map[string]interface{}{
				"code": "UNAUTHENTICATED",
			},
		}

	}
	return token, nil
}

var publicKey *ecdsa.PublicKey

func init() {
	var err error
	publicKey, err = jwt.ParseECPublicKeyFromPEM([]byte(conf.C.Auth.PubJwtkey))
	if err != nil {
		log.Fatal(err)
	}
}
