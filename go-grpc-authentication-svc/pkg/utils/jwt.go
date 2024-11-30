package utils

import (
	"errors"
	"fmt"
	"strings"

	// "fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.StandardClaims
	AccountId string
	UserRole  string
}

func (w *JwtWrapper) GenerateToken(accountId string, userRole string) (signedToken string, err error) {

	claims := &jwtClaims{
		AccountId: accountId,
		UserRole:  userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, errors.New("Couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil

}

func GetUserRolesFromToken(authHeader string) ([]string, error) {
	// Check if the Authorization header starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("invalid Authorization header format")
	}

	// Get the token string without the "Bearer " prefix
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse the token string
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		// In this example, we are using a hardcoded key for simplicity
		return []byte("r43t18sc"), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Get the "roles" claim from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// Get the "role" claim from the token
	userRole, ok := claims["UserRole"].(string)
	if !ok {
		return nil, errors.New("missing or invalid 'role' claim")
	}

	// Convert the "role" claim to a slice of strings
	rolesStr := []string{userRole}

	return rolesStr, nil
}
