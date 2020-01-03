package infrastructure

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ezio1119/fishapp-api-gateway/conf"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

func NewGraphQLHandler(resolver gen.ResolverRoot, fieldFunc graphql.FieldMiddleware) (*handler.Server, http.HandlerFunc) {
	srv := handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: resolver}))
	srv.AroundFields(fieldFunc)
	return srv, playground.Handler("GraphQL playground", conf.C.Graphql.Endpoint)
}
