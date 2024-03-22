package domain

import "time"

type User struct {
	ID         uint64     `gorm:"primary_key;column:id;<-:create"`
	Name       Name      `gorm:"embedded"`
	Password   string    `gorm:"column:password"`
	Created_At time.Time `gorm:"column:created_at;autoCreateTime:true;<-:create"`
	Updated_At time.Time `gorm:"column:updated_at;autoUpdateTime:true"`
}

func (u *User) TableName() string {
	return "users"
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}
