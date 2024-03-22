package tests

import (
	"fmt"
	"golang-gorm/app/enum"
	"golang-gorm/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteSQL(t *testing.T) {
	queryInsert := "INSERT INTO todos(content,status) VALUES(?,?)"
	result := DB.Exec(queryInsert, "Makan siang", string(enum.ACTIVE))
	helper.SaveToLogInfo(result.RowsAffected)
	assert.Nil(t, result.Error)

	result = DB.Exec(queryInsert, "Makan malam", string(enum.ACTIVE))
	helper.SaveToLogInfo(result.RowsAffected)
	assert.Nil(t, result.Error)

	result = DB.Exec(queryInsert, "Tidur", string(enum.ACTIVE))
	helper.SaveToLogInfo(result.RowsAffected)
	assert.Nil(t, result.Error)

	result = DB.Exec(queryInsert, "Belajar", string(enum.ACTIVE))
	helper.SaveToLogInfo(result.RowsAffected)
	assert.Nil(t, result.Error)
}

func TestRawSQL(t *testing.T) {
	
}

func TestEnum(t *testing.T) {
	if enum.ACTIVE == "active" {
		fmt.Println(enum.ACTIVE)
	}
}