package jwt

import (
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

// ErrInvalidToken is returned when a token fails validation.
var ErrInvalidToken = errors.New("invalid or expired token")

// Claims embeds the registered JWT claims plus the authenticated user id.
type Claims struct {
	UserID string `json:"user_id"`
	jwtlib.RegisteredClaims
}

// Generate creates a signed JWT for the given user id, valid for 24 hours.
func Generate(userID, secret string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwtlib.NewNumericDate(time.Now()),
		},
	}
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Parse validates a token string and returns the embedded claims.
func Parse(tokenString, secret string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwtlib.ParseWithClaims(tokenString, claims, func(t *jwtlib.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
