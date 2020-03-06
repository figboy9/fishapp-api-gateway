package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/graph"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure/middleware"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()
	middLe := middleware.InitMiddleware()
	r := graph.NewResolver(infrastructure.NewGrpcClient())
	srv, playground := infrastructure.NewGraphQLHandler(r, middLe)
	if conf.C.Sv.Debug {
		http.Handle(conf.C.Graphql.Playground, playground)
	}
	http.Handle(conf.C.Graphql.Endpoint, middLe.GetTokenFromReq(srv))
	http.HandleFunc("/healthy", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "healthy")
	})
	log.Fatal(http.ListenAndServe(":"+conf.C.Sv.Port, nil))
}
