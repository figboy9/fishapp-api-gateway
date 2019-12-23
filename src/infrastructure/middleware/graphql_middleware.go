package middleware

import (
	"context"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-api-gateway/conf"
)

func FieldMiddleware(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	// fmt.Printf("\nctxだよ: %#v\n", ctx.Value(reqToken{}))
	// fmt.Printf("\nFieldMiddleware: %#v\n", graphql.GetFieldContext(ctx).Object)
	gqlgenCtx := graphql.GetFieldContext(ctx)
	path := gqlgenCtx.Path()
	isMethod := gqlgenCtx.IsMethod
	if path[0] == "createPost" && isMethod {
		token := ctx.Value(reqToken{}).(string)
		userID, err := validateToken(token)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, "userID", userID)
	}
	// fmt.Printf("%#v", path.(type))

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
