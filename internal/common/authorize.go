package common

import (
	"encoding/base32"
	"github.com/BryceSun/beacon_academy/internal/account/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GetKey(t *jwt.Token) (interface{}, error) {
	key, ok := t.Header["key"].(string)
	if !ok {
		return []byte("wjwishandsome"), nil
	}
	key = base32.StdEncoding.EncodeToString([]byte(key))
	return []byte(key), nil
}

func GetToken(user *model.UserAccount) (string, error) {

	expireTime := time.Now().Add(5 * time.Minute).Unix()

	claims := make(jwt.MapClaims)
	claims["exp"] = expireTime
	claims["iat"] = time.Now().Unix()
	claims["uid"] = user.Id
	claims["unm"] = user.Username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["key"] = user.Username
	key, _ := GetKey(token)
	return token.SignedString(key)
}
