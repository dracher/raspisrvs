package middleware

import (
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

var jwtKey []byte

func init() {
	jwtKey = genJwtKey(14)
}

func genJwtKey(keyLen int) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	hmacKey := make([]byte, keyLen)
	r.Read(hmacKey)
	return hmacKey
}

func genJwtToken() string {
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 10,
		Issuer:    "dracher's pi",
		Subject:   "api_v1_token",
		Id:        "dracher",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(jwtKey)
	return ss
}

// APIJwt01 is
var APIJwt01 = jwtmiddleware.New(jwtmiddleware.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
