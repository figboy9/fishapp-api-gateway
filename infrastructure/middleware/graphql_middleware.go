package middleware

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-api-gateway/conf"

	"github.com/ezio1119/fishapp-api-gateway/graph/gqlerr"
	"github.com/ezio1119/fishapp-api-gateway/graph/model"
)

func (*middleware) Authentication(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
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

func getClaimsFromToken(t string) (*model.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(t, &model.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if c, ok := token.Claims.(*model.JwtClaims); ok && token.Valid {
		return c, nil
	}
	return nil, err
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
		log.Fatal(err)
	}
}
