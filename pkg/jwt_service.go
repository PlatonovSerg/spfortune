package pkg

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret     []byte
	algorithm  *jwt.SigningMethodHMAC
	issuer     string
	expiration time.Duration
}

func NewJWTService(secret string, issuer string, expiration time.Duration) *JWTService {
	return &JWTService{
		secret:     []byte(secret),
		algorithm:  jwt.SigningMethodHS256,
		issuer:     issuer,
		expiration: expiration,
	}
}

func (j *JWTService) GenerateToken(userID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    j.issuer,
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiration)),
	}
	token := jwt.NewWithClaims(j.algorithm, claims)
	return token.SignedString(j.secret)
}

func (j *JWTService) DecodeToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

func (j *JWTService) RetrieveTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing authentication token")
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header format")
	}
	return parts[1], nil
}
