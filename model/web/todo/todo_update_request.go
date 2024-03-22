package todo

import "golang-gorm/app/enum"

type TodoUpdateRequest struct {
	ID int64 `validate:"required,number" json:"id"`
	Content string `validate:"required,max=200,min=3" json:"content"`
	Status enum.Status `validate:"required" json:"status"`
}