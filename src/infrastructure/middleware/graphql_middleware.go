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

func (*middleware) FieldMiddleware(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	gqlgenCtx := graphql.GetFieldContext(ctx)
	if isMethod := gqlgenCtx.IsMethod; !isMethod {
		return next(ctx)
	}

	switch path := gqlgenCtx.Path().String(); path {
	// 検証せずにトークンをコンテキストに保存
	case "updateUser", "deleteUser", "refreshIDToken", "logout":
		token, err := getTokenCtx(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, resolver.JwtTokenKey, token)
	// トークンを検証してclaimsをコンテキストに保存
	case "createPost", "updatePost", "deletePost", "updateProfile", "createEntryPost", "deleteEntryPost", "createChatRoom", "sendMessage", "messageAdded", "addMemberToChatRoom":
		token, err := getTokenCtx(ctx)
		if err != nil {
			return nil, err
		}
		jwtClaims, err := validateToken(token)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, resolver.JwtClaimsCtxKey, jwtClaims)
		// トークンを検証してclaimsとトークンをコンテキストに保存
	case "deleteUserProfile":
		token, err := getTokenCtx(ctx)
		if err != nil {
			return nil, err
		}
		jwtClaims, err := validateToken(token)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, resolver.JwtTokenKey, token)
		ctx = context.WithValue(ctx, resolver.JwtClaimsCtxKey, jwtClaims)
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

	return resolver.JwtClaims{
		UserID:    userID,
		Jti:       claims.Id,
		ExpiresAt: claims.ExpiresAt,
	}, nil
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
