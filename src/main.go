package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure/middleware"
	"github.com/ezio1119/fishapp-api-gateway/registry"
)

func main() {
	conf.Readconf()
	// mux := http.NewServeMux()
	postClient := infrastructure.NewGrpcClient()
	t := time.Duration(conf.C.Sv.Timeout) * time.Second
	resolver := registry.NewGraphQLResolver(t, postClient)
	query, playground := infrastructure.NewGraphQLHandler(resolver)
	if conf.C.Sv.Debug {
		http.Handle("/graphql/playground", playground)
	}
	http.Handle("/graphql", middleware.GetToken(query))
	log.Fatal(http.ListenAndServe(":"+conf.C.Sv.Port, nil))
}
