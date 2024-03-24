package domain

import (
	"golang-gorm/app/enum"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID         uint64         `gorm:"primary_key;column:id;autoIncrement"`
	UserId     uint64         `gorm:"column:user_id"`
	Content    string         `gorm:"column:content"`
	Status     enum.Status    `gorm:"column:status"`
	Created_At time.Time      `gorm:"column:created_at;autoCreateTime:true;<-:create"`
	Updated_At time.Time      `gorm:"column:updated_at;autoUpdateTime:true"`
	Deleted_At gorm.DeletedAt `gorm:"column:deleted_at"`
	User       User           `gorm:"foreignKey:user_id;references:id"`
}

func (t *Todo) TableName() string {
	return "todos"
}
