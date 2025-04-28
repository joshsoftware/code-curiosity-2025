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

// Generates jwt token
func GenerateJWT(userId int, isAdmin bool) (string, error) {
	appCfg := config.GetAppConfig()

	claims := Claims{
		UserId:  userId,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "BeMyRoomie",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(appCfg.JWTSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Verfies JWT token
func ParseJWT(tokenStr string) (*Claims, error) {
	appCfg := config.GetAppConfig()
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return appCfg.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err

	}
}
