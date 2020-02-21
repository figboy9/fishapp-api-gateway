package repository

import (
	"context"
	"io"

	"github.com/ezio1119/fishapp-api-gateway/domain/entry_post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/usecase/repository"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type entryPostRepository struct {
	client entry_post_grpc.EntryServiceClient
}

func NewEntryRepository(c entry_post_grpc.EntryServiceClient) repository.EntryPostRepository {
	return &entryPostRepository{c}
}

func (r *entryPostRepository) Create(ctx context.Context, req *entry_post_grpc.CreateReq) (*entry_post_grpc.Entry, error) {
	return r.client.Create(ctx, req)
}

func (r *entryPostRepository) Delete(ctx context.Context, req *entry_post_grpc.DeleteReq) (*wrappers.BoolValue, error) {
	return r.client.Delete(ctx, req)
}

func (r *entryPostRepository) GetListByPostID(ctx context.Context, req *entry_post_grpc.ID) ([]*entry_post_grpc.Entry, error) {
	stream, err := r.client.GetListByPostID(ctx, req)
	if err != nil {
		return nil, err
	}
	var entries []*entry_post_grpc.Entry
	for {
		e, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}
