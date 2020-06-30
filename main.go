package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/graph"
	"github.com/ezio1119/fishapp-api-gateway/graph/dataloader"
	graphMiddle "github.com/ezio1119/fishapp-api-gateway/graph/middleware"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure"
	"github.com/ezio1119/fishapp-api-gateway/infrastructure/middleware"
	"github.com/ezio1119/fishapp-api-gateway/pb"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, conf.C.API.PostURL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	postC := pb.NewPostServiceClient(conn)

	conn, err = grpc.DialContext(ctx, conf.C.API.UserURL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	authC := pb.NewUserServiceClient(conn)

	conn, err = grpc.DialContext(ctx, conf.C.API.ChatURL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	chatC := pb.NewChatServiceClient(conn)

	conn, err = grpc.DialContext(ctx, conf.C.API.ImageURL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	imageC := pb.NewImageServiceClient(conn)

	natsConn, err := infrastructure.NewNatsStreamingConn()
	if err != nil {
		panic(err)
	}
	defer natsConn.Close()

	resolver := graph.NewResolver(postC, authC, chatC, imageC, natsConn)
	gqlMW := graphMiddle.NewGraphQLMiddleware(chatC, postC)
	gqlHandler := infrastructure.NewGraphQLHandler(resolver, gqlMW)

	if conf.C.Sv.Debug {
		http.Handle(conf.C.Graphql.Playground, infrastructure.NewPlayGroundHandler())
	}

	http.Handle(
		conf.C.Graphql.Endpoint,
		middleware.Cors(
			middleware.GetTokenFromHeader(
				dataloader.LoaderMiddleware(
					gqlHandler,
					postC,
				))))

	log.Fatal(http.ListenAndServe(":"+conf.C.Sv.Port, nil))
}
