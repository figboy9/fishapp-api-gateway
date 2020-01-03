package infrastructure

import (
	"log"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"google.golang.org/grpc"
)

func NewGrpcClient() (post_grpc.PostServiceClient, auth_grpc.AuthServiceClient) {
	conn, err := grpc.Dial(conf.C.Grpc.PostURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	postClient := post_grpc.NewPostServiceClient(conn)

	conn, err = grpc.Dial(conf.C.Grpc.AuthURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	authClient := auth_grpc.NewAuthServiceClient(conn)

	return postClient, authClient
}
