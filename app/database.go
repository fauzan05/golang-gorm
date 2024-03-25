package app

import (
	"golang-gorm/helper"
	"time"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection(log *logrus.Logger) *gorm.DB {	
    // Membuat koneksi database
    db_conf := helper.GetDatabaseConfigTest()
    dialect := mysql.Open(db_conf.Username + ":@tcp(" + db_conf.Host + ":" + db_conf.Port + ")/" + db_conf.Name + "?" + "charset=utf8mb4&parseTime=True&loc=Local")
    db, err := gorm.Open(dialect, &gorm.Config{
        Logger:     logger.New(
			&logrusWriter{Logger: log},
			logger.Config{
				SlowThreshold:             time.Second * 5,   // Slow SQL threshold
				LogLevel:                  logger.Info,  // Log level
				IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,         // Don't include params in the SQL log
				Colorful:                  false,        // Disable color
			}),
        PrepareStmt: true,
    })
    helper.HandleErrorWithPanic(err)
	// db.Logger = logger.Default.LogMode(logger.Info) // jika ingin log di terminal
    return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Infof(message, args...)
	helper.SaveToLogInfo(args)
}

