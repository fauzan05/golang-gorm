package tests

import (
	"golang-gorm/app"
	"golang-gorm/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

var DB = app.OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, DB)
	helper.SaveToLogInfo("Koneksi Berhasil")
}