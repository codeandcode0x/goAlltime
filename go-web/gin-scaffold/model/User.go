package model

import "gin-scaffold/db"

// user entity
type User struct {
	ID           uint64 `json:"id,omitempty"`
	TeamId       int64  `json:"teamid,omitempty"`
	TeamName     string `json:"teamname,omitempty" gorm:"type:varchar(300)"`
	JobId        uint64 `json:"jobid,omitempty"`
	App          string `json:"app,omitempty" gorm:"type:varchar(300)"`
	AttackTime   int64  `json:"attacktime,omitempty"`
	ResourceName string `json:"resourcename,omitempty" gorm:"type:varchar(300)"`
	AttackEvent  string `json:"attackevent,omitempty" gorm:"type:varchar(1000)"`
	UserType     int8   `json:"usertype,omitempty" gorm:"default:1"` // 1:未处理 2:已处理 3:已忽略
	BaseModel
}

// user DAO
type UserDAO interface {
	BaseDAO
	FindByID(uid uint64) (*User, error)
	FindAllByCount(count int) ([]User, error)
}

func (u *User) FindByID(uid uint64) (*User, error) {
	var instance *User
	result := db.Conn.Model(&User{}).First(&instance, uid)
	return instance, result.Error
}

func (u *User) FindAllByCount(count int) ([]User, error) {
	var users []User
	result := db.Conn.Model(&User{}).Limit(count).Order("id").Find(&users)
	return users, result.Error
}
