//go:generate go run github.com/vektah/dataloaden ProfileLoader int64 *github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc.Profile
//go:generate go run github.com/vektah/dataloaden ApplyPostLoader int64 []*github.com/ezio1119/fishapp-api-gateway/grpc/post_grpc.ApplyPost

package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/grpc/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
)

const loadersKey = "dataloaders"

type loaders struct {
	ProfileByUserID     *ProfileLoader
	ApplyPostsByPostIDs *ApplyPostLoader
}

func LoaderMiddleware(
	pC profile_grpc.ProfileServiceClient,
	aC post_grpc.PostServiceClient,
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}

		ldrs.ProfileByUserID = &ProfileLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(userIDs []int64) ([]*profile_grpc.Profile, []error) {
				res, err := pC.BatchGetProfiles(context.Background(), &profile_grpc.BatchGetProfilesReq{
					UserIds: userIDs,
				})
				if err != nil {
					return nil, []error{err}
				}
				return res.Profiles, nil
			},
		}

		ldrs.ApplyPostsByPostIDs = &ApplyPostLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(postIDs []int64) ([][]*post_grpc.ApplyPost, []error) {
				res, err := aC.BatchGetApplyPostsByPostIDs(context.Background(), &post_grpc.BatchGetApplyPostsByPostIDsReq{
					PostIds: postIDs,
				})
				if err != nil {
					return nil, []error{err}
				}
				list := make([][]*post_grpc.ApplyPost, len(postIDs))
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
