package infrastructure

import (
	"log"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
	"google.golang.org/grpc"
)

func NewGrpcClient() (post_grpc.PostServiceClient, user_grpc.UserServiceClient) {
	postConn, err := grpc.Dial(conf.C.Grpc.PostURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	postClient := post_grpc.NewPostServiceClient(postConn)

	userConn, err := grpc.Dial(conf.C.Grpc.UserURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	userClient := user_grpc.NewUserServiceClient(userConn)

	return postClient, userClient
}
