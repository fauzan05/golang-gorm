package todo

import "golang-gorm/app/enum"

type TodoResponse struct {
	Id string `json:"id"`
	Content string `json:"content"`
	Status enum.Status `json:"status"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}