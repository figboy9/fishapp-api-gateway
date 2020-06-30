package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ezio1119/fishapp-api-gateway/pb"
)

func UnmarshalSex(v interface{}) (pb.Sex, error) {
	str, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("enums must be strings")
	}
	switch str {
	case "FEMALE":
		return pb.Sex_FEMALE, nil
	case "MALE":
		return pb.Sex_MALE, nil
	default:
		return 0, fmt.Errorf("failed UnmarshalSex: recieved str: %s", str)
	}
}

func MarshalSex(o pb.Sex) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(o.String()))
	})
}

func UnmarshalOwnerType(v interface{}) (pb.OwnerType, error) {
	str, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("enums must be strings")
	}
	switch str {
	case "POST":
		return pb.OwnerType_POST, nil
	case "USER":
		return pb.OwnerType_USER, nil
	case "MESSAGE":
		return pb.OwnerType_MESSAGE, nil
	default:
		return 0, fmt.Errorf("failed UnmarshalOwnerType: recieved str: %s", str)
	}
}

func MarshalOwnerType(o pb.OwnerType) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(o.String()))
	})
}

func UnmarshalPostSortBy(v interface{}) (pb.ListPostsReq_Filter_SortBy, error) {
	str, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("enums must be strings")
	}
	switch str {
	case "CREATED_AT":
		return pb.ListPostsReq_Filter_CREATED_AT, nil
	case "MEETING_AT":
		return pb.ListPostsReq_Filter_MEETING_AT, nil
	default:
		return 0, fmt.Errorf("failed UnmarshalPostSortBy: recieved str: %s", str)
	}
}

func MarshalPostSortBy(s pb.ListPostsReq_Filter_SortBy) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, s.String())
	})
}

func UnmarshalPostOrderBy(v interface{}) (pb.ListPostsReq_Filter_OrderBy, error) {
	str, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("enums must be strings")
	}
	switch str {
	case "DESC":
		return pb.ListPostsReq_Filter_DESC, nil
	case "ASC":
		return pb.ListPostsReq_Filter_ASC, nil
	default:
		return 0, fmt.Errorf("failed UnmarshalPostOrderBy: recieved str: %s", str)
	}
}

func MarshalPostOrderBy(s pb.ListPostsReq_Filter_OrderBy) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, s.String())
	})
}
