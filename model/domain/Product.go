package domain

import "time"

type Product struct {
	ID           uint64    `gorm:"primary_key;column:id;autoIncrement"`
	Name         string    `gorm:"column:name"`
	Price        int64     `gorm:"column:price"`
	Created_At   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	Updated_At   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	LikedByUsers []User    `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:product_id;references:id;joinReferences:user_id"`
}

func (p *Product) TableName() string {
	return "products"
}
