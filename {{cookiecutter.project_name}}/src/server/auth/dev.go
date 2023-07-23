package auth

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3filter"
)

func DevJWTValidate(c context.Context, o *openapi3filter.AuthenticationInput) error {
	return nil
}