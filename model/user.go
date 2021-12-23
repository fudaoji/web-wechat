package model

const TABLE = "user"

// BaseUser 密码登录基础用户
type User struct {
	Username string `json:"username" gorm:"type:varchar(30);comment:账号"`
	Appkey   string `json:"appkey" gorm:"type:varchar(32);comment:appKey"`
	Deadline int    `json:"deadline" gorm:"type:int(10);comment:到期时间"`
}

// Verify 验证密码
func (u *User) findByUsername() *User {
	db.Table(TABLE).Where("username = ?", u.Username).First(u)
	return u
}
