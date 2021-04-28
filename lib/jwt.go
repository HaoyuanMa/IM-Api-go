package lib

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
)

var jwtKey = []byte("jsldkjfeljsflskjf")

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username,omitempty"`
}

func CreateToken(userName string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{Username: userName})
	var tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		return ""
	}
	return tokenString
}

func ParserToken(tokenHeader string) (string, bool) {
	checkToken := strings.Split(tokenHeader, " ")

	if len(checkToken) != 2 || checkToken[0] != "Bearer" {
		return "", false
	}
	var token, err = jwt.ParseWithClaims(checkToken[1], &MyClaims{}, func(tk *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || token == nil {
		return "", false
	}
	var claims, ok = token.Claims.(*MyClaims)
	if ok && token.Valid {
		return claims.Username, true
	}
	return "", false
}
