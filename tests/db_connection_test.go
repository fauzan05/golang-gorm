package tests

import (
	"golang-gorm/app"
	"golang-gorm/helper"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var DB = app.OpenConnection(&logrus.Logger{})

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, DB)
	helper.SaveToLogInfo("Koneksi Berhasil")
}