package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joshsoftware/code-curiosity-2025/internal/config"
)

type Claims struct {
	UserId  int
	IsAdmin bool
	jwt.RegisteredClaims
}

func GenerateJWT(userId int, isAdmin bool, appCfg config.AppConfig) (string, error) {
	claims := Claims{
		UserId:  userId,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(appCfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenStr string, appCfg config.AppConfig) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(appCfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
