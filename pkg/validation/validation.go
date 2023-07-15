package validation

import (
	"fmt"
	"time"

	"github.com/dhuki/go-date/config"
	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -destination=mocks/mock_validation.go -package=mocks github.com/dhuki/go-date/pkg/validation Validation

type Validation interface {
	GenerateJWTAccessToken(userID uint64) (token string, err error)
	ParseJWTAccessToken(tokenString string) (token *jwt.Token, err error)
}

type ValidationImpl struct{}

func NewValidation() Validation {
	return ValidationImpl{}
}

func (v ValidationImpl) GenerateJWTAccessToken(userID uint64) (token string, err error) {
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        fmt.Sprint(userID),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	})

	// sign the generated key using secretKey
	token, err = tokenString.SignedString([]byte(config.Conf.JwtSecret))
	return
}

func (v ValidationImpl) ParseJWTAccessToken(tokenString string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: %v", ErrSigningMethod, token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Conf.JwtSecret), nil
	})
	if err != nil {
		return
	}
	return
}
