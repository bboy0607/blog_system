package model

type Membership struct {
	*Model
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	State    uint   `json:"state"`
}

// 設定membership表名
func (m Membership) TableName() string {
	return "membership"
}
