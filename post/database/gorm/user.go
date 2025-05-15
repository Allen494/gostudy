package database

import (
	"post/database/model"
)

// 注册新用户
func RegistUser(name, password string) error {
	user := new(model.User)
	user.Name = name
	user.PassWord = password

	return nil
}
