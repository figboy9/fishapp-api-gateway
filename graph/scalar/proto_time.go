package scalar

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTimeProto(protoT timestamp.Timestamp) graphql.Marshaler {
	t, _ := ptypes.Timestamp(&protoT)
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.In(time.Local).Format(time.RFC3339)))
	})
}

func UnmarshalTimeProto(v interface{}) (timestamp.Timestamp, error) {
	if tmpStr, ok := v.(string); ok {
		t, err := time.Parse(time.RFC3339, tmpStr)
		if err != nil {
			return timestamp.Timestamp{}, err
		}
		a, _ := ptypes.TimestampProto(t)
		return *a, nil
	}
	return timestamp.Timestamp{}, errors.New("time should be RFC3339 formatted string")
}
