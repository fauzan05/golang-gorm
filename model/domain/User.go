package domain

import "time"

type User struct {
	ID           uint64    `gorm:"primary_key;column:id;autoIncrement"`
	Name         Name      `gorm:"embedded"`
	Password     string    `gorm:"column:password"`
	Created_At   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	Updated_At   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Wallet       Wallet    `gorm:"foreignKey:user_id;references:id"`
	Todos        []Todo    `gorm:"foreignKey:user_id;references:id"`
	LikeProducts []Product `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:product_id"`
}

func (u *User) TableName() string {
	return "users"
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}
