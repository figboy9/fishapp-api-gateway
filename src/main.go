package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure"
	"github.com/ezio1119/fishapp-api-gateway/middleware"
	"github.com/ezio1119/fishapp-api-gateway/registry"
)

func main() {
	postClient, userClient := infrastructure.NewGrpcClient()
	t := time.Duration(conf.C.Sv.Timeout) * time.Second
	resolver := registry.NewGraphQLResolver(t, postClient, userClient)
	srv, playground := infrastructure.NewGraphQLHandler(resolver, middleware.FieldMiddleware)
	if conf.C.Sv.Debug {
		http.Handle("/graphql/playground", playground)
	}
	http.Handle("/graphql", middleware.GetTokenFromReq(srv))
	log.Fatal(http.ListenAndServe(":"+conf.C.Sv.Port, nil))
}
