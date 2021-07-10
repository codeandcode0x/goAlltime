package model

import (
	"database/sql"
	"gin-scaffold/db"
	"time"
)

// instance entity
type User struct {
	ID           uint64         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty" gorm:"type:varchar(255)"`
	Password     string         `json:"password,omitempty" gorm:"type:varchar(1000)"`
	Email        string         `json:"email,omitempty" gorm:"type:varchar(255)"`
	Age          int            `json:"age,omitempty"`
	Birthday     time.Time      `json:"birthday,omitempty"`
	MemberNumber sql.NullString `json:"member_number,omitempty" gorm:"type:varchar(255)"`
	Role         string         `json:"Role,omitempty" gorm:"type:varchar(100)"`
	BaseModel
}

// user DAO
type UserDAO interface {
	BaseDAO
	FindAllByCount(count int) ([]User, error)
}

func (u *User) FindAllByCount(count int) ([]User, error) {
	var users []User
	result := db.Conn.Model(&User{}).Limit(count).Order("id").Find(&users)
	return users, result.Error
}
