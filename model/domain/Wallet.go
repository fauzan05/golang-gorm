package domain

import "time"

type Wallet struct {
	ID         uint64    `gorm:"primary_key;column:id;autoIncrement"`
	UserId     uint64    `gorm:"column:user_id"`
	Balance    int64     `gorm:"column:balance"`
	Created_At time.Time `gorm:"column:created_at;autoCreateTime:true;<-:create"`
	Updated_At time.Time `gorm:"column:updated_at;autoUpdateTime:true"`
	User       *User     `gorm:"foreignKey:user_id;references:id"`
}

func (w *Wallet) TableName() string {
	return "wallets"
}

