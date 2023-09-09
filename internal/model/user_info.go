package model

import (
	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	*Model
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
}

func (u UserInfo) TableName() string {
	return "user_info"
}

func (u UserInfo) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u UserInfo) Get(db *gorm.DB) (*UserInfo, error) {
	var err error
	userInfo := UserInfo{}

	// 構建查詢條件
	query := db.Model(&UserInfo{})
	if u.UserID != "" {
		query = query.Where("user_id = ?", u.UserID)
	}
	query = query.Where("is_del = ?", 0)

	// 執行查詢並將結果存儲到 userInfo
	if err = query.First(&userInfo).Error; err != nil {
		return nil, err
	}

	return &userInfo, nil
}
