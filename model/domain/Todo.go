package domain

import (
	"golang-gorm/app/enum"
	"time"
)

type Todo struct {
	ID         uint64       `gorm:"primary_key;column:id;<-:create"`
	Content    string      `gorm:"column:content"`
	Status     enum.Status `gorm:"column:status"`
	Created_At time.Time   `gorm:"column:created_at;autoCreateTime:true;<-:create"`
	Updated_At time.Time   `gorm:"column:updated_at;autoUpdateTime:true"`
}

func (t *Todo) TableName() string {
	return "todos"
}