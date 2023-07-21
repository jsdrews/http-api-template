package auth

import (
	"context"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"google.golang.org/api/idtoken"
)

func GoogleJWTValidate(c context.Context, input *openapi3filter.AuthenticationInput) error {
	bearerTokenString := input.RequestValidationInput.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerTokenString, "Bearer ") // Bearer <token>
	tokenString := splitToken[1]
	_, err := idtoken.Validate(context.Background(), tokenString, "")
	return err
}
