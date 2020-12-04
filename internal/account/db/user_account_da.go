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