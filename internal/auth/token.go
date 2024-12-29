package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GetBearerToken(headers http.Header) (string, error) {
	authInfo := headers.Get("Authorization")
	tokenParts := strings.Split(authInfo, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", fmt.Errorf("no bearer token found")
	}
	return tokenParts[1], nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	jwtClaims, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("invalid token claims")
	}

	if time.Now().After(jwtClaims.ExpiresAt.Time) {
		return uuid.UUID{}, fmt.Errorf("jwt token expired")
	}

	userID, err := uuid.Parse(jwtClaims.Subject)
	if err != nil {
		return uuid.UUID{}, err
	}
	return userID, nil
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn).UTC()),
		Subject:   userID.String(),
	})
	signedToken, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
