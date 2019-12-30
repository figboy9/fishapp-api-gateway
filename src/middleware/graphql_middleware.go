package middleware

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver"
)

func FieldMiddleware(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	gqlgenCtx := graphql.GetFieldContext(ctx)
	if isMethod := gqlgenCtx.IsMethod; !isMethod {
		return next(ctx)
	}
	switch path := gqlgenCtx.Path(); path[0] {
	case "createPost", "updatePost", "deletePost", "updateUser", "deleteUser", "refreshToken", "logout":
		token, err := getTokenCtx(ctx)
		if err != nil {
			return nil, err
		}
		jwtClaimsCtx, err := validateToken(token)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, resolver.JwtCtxKey, jwtClaimsCtx)
	}
	return next(ctx)
}

// トークンからclaimsを取り出し、resolver用のコンテキストの構造体に入れる
func validateToken(t string) (resolver.JwtClaims, error) {
	var claims jwt.StandardClaims
	_, err := jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return resolver.JwtClaims{}, err
	}
	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return resolver.JwtClaims{}, err
	}
	jwtClaimsCtx := resolver.JwtClaims{
		UserID:    userID,
		Jti:       claims.Id,
		ExpiresAt: claims.ExpiresAt,
	}

	return jwtClaimsCtx, nil
}

func getTokenCtx(ctx context.Context) (string, error) {
	v := ctx.Value(tokenCtxKey)
	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("token not found")
	}
	return token, nil
}

var publicKey *ecdsa.PublicKey

func init() {
	var err error
	data := []byte(conf.C.Auth.PubJwtkey)
	if conf.C.Sv.Debug {
		// 開発環境はpemから読み込む
		data, err = ioutil.ReadFile("./dev_pub_jwtkey.pem")
		if err != nil {
			log.Fatal(err)
		}
	}
	publicKey, err = jwt.ParseECPublicKeyFromPEM(data)
	if err != nil {
		log.Fatal(err)
	}
}
