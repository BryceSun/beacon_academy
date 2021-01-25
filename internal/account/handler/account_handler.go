package handler

import (
	"crypto/md5"
	"fmt"
	. "github.com/BryceSun/beacon_academy/init_config"
	"github.com/BryceSun/beacon_academy/internal/account/db"
	"github.com/BryceSun/beacon_academy/internal/account/model"
	. "github.com/BryceSun/beacon_academy/internal/common"
	"time"
)

func AddUserAccount(u *model.UserAccount) (int64, error) {
	u.Password = fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))
	return db.AddUserAccount(u)
}

func AddUserAccountByEmail(u *model.UserAccount) (string, error) {
	to := []string{
		u.Email,
	}
	msg := "From: 15818366547@163.com\r\n" +
		"To: 865874804@qq.com\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n"
	err := Eamil.SendMail(to, msg)
	if err != nil {
		return "", err
	}
	return "ok", err
}

//todo
func AddUserAccountByPhone(u *model.UserAccount) (int64, error) {
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
	token, err := GetToken(du)
	if err == nil {
		Redis.Set(TokenKey(du.Id), token, time.Minute*5)
	}
	return token, err
}

func DeleteUserToken(uid int64) (string, error) {
	tokenKey := TokenKey(uid)
	val, err := Redis.Del(tokenKey).Result()
	if err != nil {
		return "", err
	}
	if val == 0 {
		return "", &OperateFailure
	}
	return "logout!", nil
}
