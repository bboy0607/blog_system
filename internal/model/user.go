package model

import (
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

func (u User) ResetUserPassword(db *gorm.DB) error {
	db = db.Model(&User{}).Where("username = ? AND is_del = ?", u.Username, 0)
	return db.Update(&u).Error
}
