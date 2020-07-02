package infrastructure

import (
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ezio1119/fishapp-api-gateway/conf"
	gen "github.com/ezio1119/fishapp-api-gateway/graph/generated"
	"github.com/ezio1119/fishapp-api-gateway/graph/gqlerr"
	"github.com/ezio1119/fishapp-api-gateway/graph/middleware"
)

func NewGraphQLHandler(r gen.ResolverRoot, gqlMW middleware.GraphQLMiddleware) *handler.Server {
	c := gen.Config{Resolvers: r}

	c.Directives.IsAuthenticated = gqlMW.IsAuthenticated
	c.Directives.IsMember = gqlMW.IsMember
	c.Directives.IsPostOwner = gqlMW.IsPostOwner
	c.Directives.IsApplyPostOwner = gqlMW.IsApplyPostOwner

	srv := handler.New(gen.NewExecutableSchema(c))

	var f transport.WebsocketInitFunc = middleware.GetTokenFromWebsocketInit

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

	return srv
}

func NewPlayGroundHandler() http.Handler {
	p := playground.Handler("GraphQL playground", conf.C.Graphql.URL)
	if !conf.C.Sv.Debug {
		return basicAuthMW(p)
	}

	return p
}

func basicAuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || len(strings.TrimSpace(u)) < 1 || len(strings.TrimSpace(p)) < 1 {
			unauthorised(w)
			return
		}

		// This is a dummy check for credentials.
		if u != conf.C.Graphql.Playground.User || p != conf.C.Graphql.Playground.Pass {
			unauthorised(w)
			return
		}

		// If required, Context could be updated to include authentication
		// related data so that it could be used in consequent steps.
		next.ServeHTTP(w, r)
	})
}

func unauthorised(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
	w.WriteHeader(http.StatusUnauthorized)
}
