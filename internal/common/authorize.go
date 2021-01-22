package common

import (
	"encoding/base32"
	"github.com/BryceSun/beacon_academy/internal/account/model"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type ClaimsPlus struct {
	Username  string
	UserId    int64
	HeaderImg string
	jwt.StandardClaims
}

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

	claims := ClaimsPlus{}
	claims.ExpiresAt = expireTime
	claims.IssuedAt = time.Now().Unix()
	claims.UserId = user.Id
	claims.Username = user.Username

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["key"] = user.Username
	key, _ := GetKey(token)
	return token.SignedString(key)
}

func UpdateToken(claims ClaimsPlus) (string, error) {
	expireTime := time.Now().Add(5 * time.Minute).Unix()
	claims.ExpiresAt = expireTime
	claims.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key, _ := GetKey(token)
	return token.SignedString(key)
}

func TokenKey(uid int64) string {
	tokenKey := "token:" + strconv.FormatInt(uid, 10)
	return tokenKey
}
