package presenter

import (
	"github.com/ezio1119/fishapp-api-gateway/domain/entry_post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/domain/graphql"
	"github.com/ezio1119/fishapp-api-gateway/domain/post_grpc"
)

type PostPresenter interface {
	TransformPostGraphQL(*post_grpc.Post) (*graphql.Post, error)
	TransformListPostGraphQL([]*post_grpc.Post) ([]*graphql.Post, error)
	TransformEntryPostGraphQL(*entry_post_grpc.Entry) (*graphql.EntryPost, error)
	TransformEntriesGraphQL([]*entry_post_grpc.Entry) ([]*graphql.EntryPost, error)
}
