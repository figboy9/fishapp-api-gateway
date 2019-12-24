package middleware

import (
	"context"
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver"
)

func FieldMiddleware(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	gqlgenCtx := graphql.GetFieldContext(ctx)
	path := gqlgenCtx.Path()
	isMethod := gqlgenCtx.IsMethod
	if (path[0] == "createPost" || path[0] == "updatePost" || path[0] == "deletePost" || path[0] == "updateUser" || path[0] == "deleteUser") && isMethod {
		token, err := getTokenCtx(ctx)
		if err != nil {
			return nil, err
		}
		userID, err := validateToken(token)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, resolver.UserIDCtxKey, userID)
	}

	return next(ctx)
}

func validateToken(t string) (int64, error) {
	jwtkey := []byte(conf.C.Auth.Jwtkey)
	var claims jwt.StandardClaims
	_, err := jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil {
		return 0, err
	}
	userID, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func getTokenCtx(ctx context.Context) (string, error) {
	v := ctx.Value(idTokenCtxKey)
	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("token not found")
	}
	return token, nil
}
