package graph

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func extractIDFromCursor(c string) (int64, error) {
	byteCursor, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return 0, err
	}
	splitCursor := strings.Split(string(byteCursor), ":")
	if splitCursor[0] != "post" && len(splitCursor) != 2 {
		return 0, &gqlerror.Error{
			Message: "wrong cursor format",
			Extensions: map[string]interface{}{
				"code": "BAD_USER_INPUT",
			},
		}
	}
	id, err := strconv.ParseInt(splitCursor[1], 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func genCursorFromID(i int64) string {
	strID := strconv.FormatInt(i, 10)
	return base64.StdEncoding.EncodeToString([]byte("post:" + strID))
}
