package db

import (
	"github.com/BryceSun/beacon_academy/init_config"
	"github.com/BryceSun/beacon_academy/internal/account/model"
)

func AddUserAccount(user *model.UserAccount) (int64, error) {
	r, err := init_config.Db.Exec(`insert into bgs.user_account(username,email,phone, password) values (?,?,?,?)`,
		user.Username, user.Email, user.Phone, user.Password)
	if err != nil {
		return 0, err
	}
	return r.LastInsertId()
}

func GetUserAccount(user *model.UserAccount) (*model.UserAccount, error) {
	row := init_config.Db.QueryRow("select username,password,phone,email,id from user_account where username = ? or phone = ? or email = ? limit 1",
		user.Username, user.Phone, user.Email)
	du := &model.UserAccount{}
	err := row.Scan(&du.Username, &du.Password, &du.Phone, &du.Email, &du.Id)
	if err != nil {
		return nil, err
	}
	return du, nil
}
