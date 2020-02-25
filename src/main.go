package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure/middleware"
	"github.com/ezio1119/fishapp-api-gateway/registry"
)

func main() {
	postC, entryPostC, authC, profileC, chatC := infrastructure.NewGrpcClient()
	t := time.Duration(conf.C.Sv.Timeout) * time.Second
	r := registry.NewRegistry(postC, entryPostC, authC, profileC, chatC, t)
	middLe := middleware.InitMiddleware()
	srv, playground := infrastructure.NewGraphQLHandler(r.NewResolver(), middLe.FieldMiddleware)
	if conf.C.Sv.Debug {
		http.Handle(conf.C.Graphql.Playground, playground)
	}
	http.Handle(conf.C.Graphql.Endpoint, middLe.GetTokenFromReq(srv))
	http.HandleFunc("/healthy", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "healthy")
	})
	log.Fatal(http.ListenAndServe(":"+conf.C.Sv.Port, nil))
}
