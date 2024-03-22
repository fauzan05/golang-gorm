package domain

import "golang-gorm/app/enum"

type Todo struct {
	Id int
	Content string
	Status enum.Status
	Created_At string
	Updated_At string
}