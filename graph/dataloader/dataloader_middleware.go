//go:generate go run github.com/vektah/dataloaden ApplyPostLoader int64 []*github.com/ezio1119/fishapp-api-gateway/pb.ApplyPost

package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/pb"
)

const loadersKey = "dataloaders"

type loaders struct {
	ApplyPostsByPostIDs *ApplyPostLoader
}

func LoaderMiddleware(
	next http.Handler,
	aC pb.PostServiceClient,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}

		ldrs.ApplyPostsByPostIDs = &ApplyPostLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(postIDs []int64) ([][]*pb.ApplyPost, []error) {
				res, err := aC.BatchGetApplyPostsByPostIDs(context.Background(), &pb.BatchGetApplyPostsByPostIDsReq{
					PostIds: postIDs,
				})
				if err != nil {
					return nil, []error{err}
				}
				list := make([][]*pb.ApplyPost, len(postIDs))
				for i, pID := range postIDs {
					for _, a := range res.ApplyPosts {
						if a.PostId == pID {
							list[i] = append(list[i], a)
						}
					}
				}
				return list, nil
			},
		}
		dlCtx := context.WithValue(r.Context(), loadersKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func For(ctx context.Context) loaders {
	return ctx.Value(loadersKey).(loaders)
}
