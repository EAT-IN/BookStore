package model

import (
	hash "BookStore/utils"
	"errors"
	"time"
)

type User struct {
	ID         int
	Username   string    `gorm:"type:varchar(100);NOT NULL;unique_index:uqi"`
	Password   string    `gorm:"type:varchar(100);NOT NULL"`
	Email      string    `gorm:"type:varchar(100);NOT NULL;index:idx"`
	CreateTime time.Time `gorm:"type:datetime;NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"type:datetime;NOT NULL;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// 根据用户名和密码查询用户
func CheckUserNameAndPassword(username, password string) (*User, error) {
	user := new(User)
	DB.Where("username=?", username).Find(user)
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	} else {
		err = hash.ValidatePassword(password, user.Password)
		if err != nil {
			return nil, errors.New("密码错误")
		} else {
			return user, nil
		}

	}

}

// 唯一索引查重
func CheckUserName(username string) bool {
	var count int
	DB.Model(&User{}).Where("username=?", username).Count(&count)
	if count > 0 {
		return false
	} else {
		return true
	}

}

// 保存用户
func SaveUser(username, password, email string) error {
	if CheckUserName(username) {
		// 进行密码加密
		passStr, _ := hash.GeneratePassword(password)

		user := &User{
			Username: username,
			Password: passStr,
			Email:    email,
		}
		DB.Create(user)
		return nil
	} else {
		return errors.New("用户已经存在")
	}

}
