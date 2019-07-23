package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(`setting.AppSetting.JwtSecret`)

type Jwt interface {
	GenerateToken() (string, error)
	ParseToken(token string) (interface{}, error)
}

type Claims struct {
	UserID primitive.ObjectID
	jwt.StandardClaims
}

type AdminJwt struct {
	UserID primitive.ObjectID
}

// GenerateToken generate tokens used for auth
func (admin *AdminJwt) GenerateToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(10 * time.Hour)

	claims := Claims{
		admin.UserID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "hjd",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func (admin *AdminJwt) ParseToken(token string) (interface{}, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
