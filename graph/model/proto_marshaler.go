package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"

	"github.com/ezio1119/fishapp-api-gateway/grpc/post_grpc"
	"github.com/ezio1119/fishapp-api-gateway/grpc/profile_grpc"
)

func UnmarshalSex(v interface{}) (profile_grpc.Sex, error) {
	str, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("enums must be strings")
	}
	switch str {
	case "FEMALE":
		return profile_grpc.Sex_FEMALE, nil
	case "MALE":
		return profile_grpc.Sex_MALE, nil
	default:
		return 0, fmt.Errorf("failed UnmarshalSex: recieved str: %s", str)
	}
}

func MarshalSex(s profile_grpc.Sex) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(s.String()))
	})
}

func UnmarshalPostSortBy(v interface{}) (post_grpc.ListPostsReq_Filter_SortBy, error) {
	str, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("enums must be strings")
	}
	switch str {
	case "CREATED_AT":
		return post_grpc.ListPostsReq_Filter_CREATED_AT, nil
	case "MEETING_AT":
		return post_grpc.ListPostsReq_Filter_MEETING_AT, nil
	default:
		return 0, fmt.Errorf("failed UnmarshalPostSortBy: recieved str: %s", str)
	}
}

func MarshalPostSortBy(s post_grpc.ListPostsReq_Filter_SortBy) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, s.String())
	})
}

func UnmarshalPostOrderBy(v interface{}) (post_grpc.ListPostsReq_Filter_OrderBy, error) {
	str, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("enums must be strings")
	}
	switch str {
	case "DESC":
		return post_grpc.ListPostsReq_Filter_DESC, nil
	case "ASC":
		return post_grpc.ListPostsReq_Filter_ASC, nil
	default:
		return 0, fmt.Errorf("failed UnmarshalPostOrderBy: recieved str: %s", str)
	}
}

func MarshalPostOrderBy(s post_grpc.ListPostsReq_Filter_OrderBy) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, s.String())
	})
}
