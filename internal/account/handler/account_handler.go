package handler

import (
	"crypto/md5"
	"fmt"
	"github.com/BryceSun/beacon_academy/internal/account/db"
	"github.com/BryceSun/beacon_academy/internal/account/model"
	. "github.com/BryceSun/beacon_academy/internal/common"
)

func AddUserAccount(u *model.UserAccount) (int64, error) {
	u.Password = fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))
	return db.AddUserAccount(u)
}

func GetUserToken(u *model.UserAccount) (string, error) {
	if len(u.Password) == 0 || len(u.Username) == 0 {
		return "", ParamsAreNeeded.New("请输入用户名和密码")
	}
	du, err := db.GetUserAccount(u)
	if err != nil {
		return "", SystemError.New(err.Error())
	}
	password := fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))
	if password != du.Password {
		return "", ParamsAreWrong.New("密码错误")
	}
	return GetToken(du)
}

func DeleteUserToken(id int) {

}
