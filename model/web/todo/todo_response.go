package todo

import (
	"golang-gorm/app/enum"
	"time"
)

type TodoResponse struct {
	ID         int64       `json:"id"`
	Content    string      `json:"content"`
	Status     enum.Status `json:"status"`
	Created_At time.Time   `json:"created_at"`
	Updated_At time.Time   `json:"updated_at"`
}
