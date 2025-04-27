package models

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"column:id; primaryKey; autoIncrement"`
	Username  string `json:"username" gorm:"column:username; type:varchar(255); not null; unique"`
	Password  string `json:"password" gorm:"column:password; type:varchar(255)"`
	Email     string `json:"email" gorm:"column:email; type:varchar(255)"`
	Role      string `json:"role" gorm:"column:role; type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Products  []Product `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "user"
}
