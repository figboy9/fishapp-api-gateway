//go:generate go run github.com/vektah/dataloaden ProfileLoader int64 *github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc.Profile
package dataloader

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
)

const loadersKey = "dataloaders"

type loaders struct {
	ProfileByUserID *ProfileLoader
}

func LoaderMiddleware(
	pC profile_grpc.ProfileServiceClient,
	next http.Handler,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}

		ldrs.ProfileByUserID = &ProfileLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(userIds []int64) ([]*profile_grpc.Profile, []error) {
				fmt.Println(userIds)
				res, err := pC.BatchGetProfiles(context.Background(), &profile_grpc.BatchGetProfilesReq{
					UserIds: userIds,
				})
				if err != nil {
					return nil, []error{err}
				}
				if len(res.Profiles) != len(userIds) {
					return nil, []error{errors.New("some of the user_ids does not have a profile")}
				}
				return res.Profiles, nil
			},
		}
		dlCtx := context.WithValue(r.Context(), loadersKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func For(ctx context.Context) loaders {
	return ctx.Value(loadersKey).(loaders)
}
