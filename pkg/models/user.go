package models

import (
	"github.com/keepchen/go-sail/v3/orm"
	"time"
)

// User 用户表
type User struct {
	orm.BaseModel
	UID         string     `gorm:"column:uid;type:varchar(150);NOT NULL;uniqueIndex:unique_uid;comment:账户唯一id"`
	Username    string     `gorm:"column:username;type:varchar(70);NOT NULL;uniqueIndex:unique_username;comment:用户名"`
	Password    string     `gorm:"column:password;type:varchar(1024);NOT NULL;comment:密码"`
	LatestLogin *time.Time `gorm:"column:latest_login;type:datetime;comment:上次登录时间"`
}

func (o *User) TableName() string {
	return "tb_users"
}
