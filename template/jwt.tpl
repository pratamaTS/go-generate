package jwt

import (
	"fmt"
	"go-generate/entity"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

var jwtSecret []byte

// SetJWTSecret is
func SetJWTSecret(secret string) {
	jwtSecret = []byte(secret)
}

func Parse(jwtString string) (jojoJWT entity.JojoJWT, isValid, isExpired bool) {
	jojoJWT.Token = jwtString
	token, err := jwtgo.Parse(jwtString, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		if len(jwtSecret) == 0 {
			jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		}
		return jwtSecret, nil
	})
	if err != nil {
		return jojoJWT, false, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if token.Valid == true {
			mapstructure.Decode(token.Claims, &jojoJWT)
			isValid = token.Valid
			isExpired = time.Now().Unix() > int64(claims["exp"].(float64))
			return jojoJWT, isValid, isExpired
		}
	}

	return jojoJWT, isValid, isExpired
}

// ExtractAuthorizationHeader is
func ExtractAuthorizationHeader(rawToken string) string {
	token := rawToken
	splitToken := strings.Split(rawToken, "bearer ")
	if len(splitToken) > 1 {
		token = splitToken[1]
	} else {
		splitToken := strings.Split(rawToken, "Bearer ")
		if len(splitToken) > 1 {
			token = splitToken[1]
		}
	}
	return token
}
