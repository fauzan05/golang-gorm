package app

import (
	"golang-gorm/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	db_conf := helper.GetDatabaseConfigTest()
	dialect := mysql.Open(db_conf.Username + ":@tcp(" + db_conf.Host + ":" + db_conf.Port + ")/" + db_conf.Name + "?" + "charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	helper.HandleErrorWithPanic(err)
	return db
}