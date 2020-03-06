package infrastructure

import (
	"log"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/grpc/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/chat_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
	"google.golang.org/grpc"
)

func NewGrpcClient() (
	post_grpc.PostServiceClient,
	auth_grpc.AuthServiceClient,
	profile_grpc.ProfileServiceClient,
	chat_grpc.ChatServiceClient,
) {
	conn, err := grpc.Dial(conf.C.API.PostURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	postC := post_grpc.NewPostServiceClient(conn)

	conn, err = grpc.Dial(conf.C.API.AuthURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	authC := auth_grpc.NewAuthServiceClient(conn)

	conn, err = grpc.Dial(conf.C.API.ProfileURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	profileC := profile_grpc.NewProfileServiceClient(conn)

	conn, err = grpc.Dial(conf.C.API.ChatURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	chatC := chat_grpc.NewChatServiceClient(conn)
	return postC, authC, profileC, chatC
}
