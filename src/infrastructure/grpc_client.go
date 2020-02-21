package infrastructure

import (
	"log"

	"github.com/ezio1119/fishapp-api-gateway/domain/chat_grpc"

	"github.com/ezio1119/fishapp-api-gateway/conf"
	"github.com/ezio1119/fishapp-api-gateway/domain/auth_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/entry_post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/profile_grpc"
	"google.golang.org/grpc"
)

func NewGrpcClient() (post_grpc.PostServiceClient, entry_post_grpc.EntryServiceClient, auth_grpc.AuthServiceClient, profile_grpc.ProfileServiceClient, chat_grpc.ChatServiceClient) {
	conn, err := grpc.Dial(conf.C.Grpc.PostURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	postClient := post_grpc.NewPostServiceClient(conn)
	entryPostClient := entry_post_grpc.NewEntryServiceClient(conn)

	conn, err = grpc.Dial(conf.C.Grpc.AuthURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	authClient := auth_grpc.NewAuthServiceClient(conn)

	conn, err = grpc.Dial(conf.C.Grpc.ProfileURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	profileClient := profile_grpc.NewProfileServiceClient(conn)

	conn, err = grpc.Dial(conf.C.Grpc.ChatURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	chatClient := chat_grpc.NewChatServiceClient(conn)

	return postClient, entryPostClient, authClient, profileClient, chatClient
}
