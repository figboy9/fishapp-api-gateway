package gqlerr

import (
	"fmt"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ForbiddenError(format string, a ...interface{}) *gqlerror.Error {
	return &gqlerror.Error{
		Message:    fmt.Sprintf(format, a...),
		Extensions: map[string]interface{}{"code": "FORBIDDEN"},
	}
}

func UserInputError(format string, a ...interface{}) *gqlerror.Error {
	return &gqlerror.Error{
		Message:    fmt.Sprintf(format, a...),
		Extensions: map[string]interface{}{"code": "BAD_USER_INPUT"},
	}
}

func AuthenticationError(format string, a ...interface{}) *gqlerror.Error {
	return &gqlerror.Error{
		Message:    fmt.Sprintf(format, a...),
		Extensions: map[string]interface{}{"code": "UNAUTHENTICATED"},
	}
}

func InternalServerError(format string, a ...interface{}) *gqlerror.Error {
	return &gqlerror.Error{
		Message:    fmt.Sprintf(format, a...),
		Extensions: map[string]interface{}{"code": "INTERNAL_SERVER_ERROR"},
	}
}
