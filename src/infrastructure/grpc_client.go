package infrastructure

import (
	"log"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"google.golang.org/grpc"
)

func NewGrpcClient() post_grpc.PostServiceClient {
	postConn, err := grpc.Dial(conf.C.Grpc.PostURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	postClient := post_grpc.NewPostServiceClient(postConn)
	return postClient
}
