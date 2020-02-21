package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/entry_post_grpc"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type EntryPostRepository interface {
	Create(ctx context.Context, req *entry_post_grpc.CreateReq) (*entry_post_grpc.Entry, error)
	Delete(ctx context.Context, req *entry_post_grpc.DeleteReq) (*wrappers.BoolValue, error)
	GetListByPostID(ctx context.Context, req *entry_post_grpc.ID) ([]*entry_post_grpc.Entry, error)
}
