package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/graph"
	"github.com/ezio1119/fishapp-api-gateway/graph/dataloader"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure/middleware"
)

func main() {

	pubsubC := infrastructure.NewCloudPubSubClient()
	defer pubsubC.Close()
	p, a, pro, c := infrastructure.NewGrpcClient()
	middLe := middleware.InitMiddleware()
	r := graph.NewResolver(p, a, pro, c, pubsubC)
	srv, playground := infrastructure.NewGraphQLHandler(r, middLe)
	if conf.C.Sv.Debug {
		http.Handle(conf.C.Graphql.Playground, playground)
	}
	http.Handle(conf.C.Graphql.Endpoint,
		dataloader.LoaderMiddleware(
			pro,
			p,
			middLe.GetTokenFromReq(srv),
		))
	http.HandleFunc("/healthy", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "healthy")
	})

	log.Fatal(http.ListenAndServe(":"+conf.C.Sv.Port, nil))
}
