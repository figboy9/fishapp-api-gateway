package infrastructure

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ezio1119/fishapp-api-gateway/interfaces/resolver"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

func NewGraphQLHandler(resolver *resolver.Resolver, fieldFunc graphql.FieldMiddleware) (*handler.Server, http.HandlerFunc) {
	srv := handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: resolver}))
	srv.AroundFields(fieldFunc)
	// srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {

	// 	fmt.Printf("\nOperationMiddleware: %#v\n", graphql.GetOperationContext(ctx))
	// 	return next(ctx)
	// })

	playground := playground.Handler("GraphQL playground", "/graphql")
	return srv, playground
}
