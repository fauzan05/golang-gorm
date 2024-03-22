package tests

import (
	"golang-gorm/app"
	"golang-gorm/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db = app.OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
	helper.SaveToLogInfo("Koneksi Berhasil")
}