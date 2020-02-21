package infrastructure

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure/middleware"
	gen "github.com/ezio1119/fishapp-api-gateway/interfaces/resolver/graphql"
)

func NewGraphQLHandler(resolver gen.ResolverRoot, fieldFunc graphql.FieldMiddleware) (*handler.Server, http.HandlerFunc) {
	srv := handler.New(gen.NewExecutableSchema(gen.Config{Resolvers: resolver}))
	var f transport.WebsocketInitFunc = middleware.GetTokenFromWebsocketInit
	srv.AddTransport(transport.Websocket{InitFunc: f, KeepAlivePingInterval: 10 * time.Second})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	srv.AroundFields(fieldFunc)
	return srv, playground.Handler("GraphQL playground", conf.C.Graphql.Endpoint)
}
