package app

import (
	"golang-gorm/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	db_conf := helper.GetDatabaseConfigProd()
	dialect := mysql.Open(db_conf.Username + ":@tcp(" + db_conf.Host + ":" + db_conf.Port + ")/" + db_conf.Name + "?" + "charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})
	helper.HandleErrorWithPanic(err)
	return db
}