package model

import (
	"fmt"

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

// 註冊用戶
func (m User) Create(db *gorm.DB) error {
	fmt.Println(m)
	return db.Create(&m).Error
}
