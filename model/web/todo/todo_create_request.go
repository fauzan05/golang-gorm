package todo

import "golang-gorm/app/enum"

type TodoCreateRequest struct {
	Content string `validate:"required,max=200,min=3" json:"content"`
	Status enum.Status `validate:"required"`
}