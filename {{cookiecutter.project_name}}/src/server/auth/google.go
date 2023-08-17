package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"google.golang.org/api/idtoken"
)

func GoogleJWTValidate(c context.Context, input *openapi3filter.AuthenticationInput) error {
	bearerTokenString := input.RequestValidationInput.Request.Header.Get("Authorization")
	if bearerTokenString == "" {
		return fmt.Errorf("no authorization header provided")
	}
	splitToken := strings.Split(bearerTokenString, "Bearer ") // Bearer <token>
	if len(splitToken) != 2 {
		return fmt.Errorf("invalid authorization header provided")
	}
	tokenString := splitToken[1]
	_, err := idtoken.Validate(context.Background(), tokenString, "")
	return err
}
