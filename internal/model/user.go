package model

import (
	"membership_system/pkg/errcode"

	"github.com/jinzhu/gorm"
)

type User struct {
	*Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	State    uint8  `json:"state"`
}

// 設定會員系統表名
func (m User) TableName() string {
	return "membership"
}

// 創建用戶資訊
func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u User) Activate(db *gorm.DB) error {
	db = db.Model(&User{}).Where("username = ? AND is_del = ?", u.Username, 0)
	return db.Update(&u).Error
}

func (u User) CheckEmail(db *gorm.DB) error {
	if u.Email != "" {
		db = db.Where("email = ?", u.Email)
	}
	err := db.First(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (u User) ResetUserPassword(db *gorm.DB) error {
	db = db.Model(&User{}).Where("email = ? AND is_del = ?", u.Email, 0)
	return db.Update(&u).Error
}

func (u User) ValidateUserCredentials(db *gorm.DB) error {
	var user User
	if u.Username != "" {
		db = db.Where("username = ?", u.Username).First(&user)
	}
	if db.Error != nil {
		return errcode.ErrorUserNotFound
	}
	//判斷使用者是否啟用
	if user.State == 0 {
		return errcode.ErrorUserNotActivated
	}
	//判斷密碼是否正確
	if user.Password != u.Password {
		return errcode.ErrorPasswordNotCorrect
	}
	return nil
}
