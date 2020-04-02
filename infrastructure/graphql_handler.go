package infrastructure

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ezio1119/fishapp-api-gateway/conf"
	gen "github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/graph/gqlerr"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure/middleware"
)

func NewGraphQLHandler(r gen.ResolverRoot, middLe middleware.Middleware) (*handler.Server, http.HandlerFunc) {
	c := gen.Config{Resolvers: r}
	c.Directives.IsAuthenticated = middLe.Authentication
	srv := handler.New(gen.NewExecutableSchema(c))
	var f transport.WebsocketInitFunc = middLe.GetTokenFromWebsocketInit
	srv.AddTransport(transport.Websocket{InitFunc: f, KeepAlivePingInterval: 10 * time.Second})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetErrorPresenter(gqlerr.ErrorPresenter())
	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return srv, playground.Handler("GraphQL playground", conf.C.Graphql.Endpoint)
}
