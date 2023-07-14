package validation

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -destination=mocks/mock_validation.go -package=mocks github.com/dhuki/go-date/pkg/validation Validation

type Validation interface {
	GenerateJWTAccessToken() (token string, err error)
}

type ValidationImpl struct{}

func NewValidation() Validation {
	return ValidationImpl{}
}

func (v ValidationImpl) GenerateJWTAccessToken() (token string, err error) {
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1) * time.Hour)),
	})

	// sign the generated key using secretKey
	token, err = tokenString.SignedString("secret")
	return
}
