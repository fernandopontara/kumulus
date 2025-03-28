package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("chave-secreta")

type Claims struct {
	UserID         string `json:"user_id"`
	OrganizationID string `json:"org_id"`
	Role           string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, orgID, role string) (string, error) {
	exp := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:         userID,
		OrganizationID: orgID,
		Role:           role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func VerifyPassword(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
