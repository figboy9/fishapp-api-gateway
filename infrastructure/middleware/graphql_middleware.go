package middleware

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/graph"
)

func (*middleware) AuthMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver, authAPI bool) (res interface{}, err error) {
	t, err := getTokenCtx(ctx)
	if err != nil {
		return nil, err
	}
	if authAPI {
		ctx = context.WithValue(ctx, graph.JwtTokenKey, t)
		return next(ctx)
	}
	c, err := getClaimsToken(t)
	if err != nil {
		return nil, err
	}
	if ctx, err = setJwtClaimsCtx(ctx, c); err != nil {
		return nil, err
	}

	return next(ctx)
}

type Claims struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
	jwt.StandardClaims
}

func setJwtClaimsCtx(ctx context.Context, c *Claims) (context.Context, error) {
	userID, err := strconv.ParseInt(c.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, graph.JwtClaimsCtxKey, graph.JwtClaims{
		UserID:    userID,
		Jti:       c.Id,
		ExpiresAt: c.ExpiresAt,
	}), nil
}

func getClaimsToken(t string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func getTokenCtx(ctx context.Context) (string, error) {
	v := ctx.Value(jwtTokenKey)
	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("token not found")
	}
	return token, nil
}

var publicKey *ecdsa.PublicKey

func init() {
	publicKey, err = jwt.ParseECPublicKeyFromPEM([]byte(conf.C.Auth.PubJwtkey))
	if err != nil {
		log.Fatal(err)
	}
}
