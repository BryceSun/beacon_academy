package handler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/BryceSun/beacon_academy/internal/account/db"
	"github.com/BryceSun/beacon_academy/internal/account/model"
	"github.com/BryceSun/beacon_academy/internal/common/utils"
)

func AddUserAccount(u *model.UserAccount) {
	u.Password = fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))
	db.AddUserAccount(u)
}

func GetUserToken(u *model.UserAccount) (string, error) {
	if len(u.Password) == 0 || len(u.Username) == 0 {
		return "", errors.New("params are needed")
	}
	du, err := db.GetUserAccount(u)
	if err != nil {
		return "", err
	}
	password := fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))
	if password != du.Password {
		return "", errors.New("password is wrong")
	}
	return utils.GetToken(du)
}
