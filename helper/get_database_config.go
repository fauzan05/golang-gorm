package helper

import (
	"strconv"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Type     string
	Name     string
	Host     string
	Port     string
	Username string
	Password string
	Pool     DatabasePool
}

type DatabasePool struct {
	Idle     int
	Max      int
	Lifetime int
}

var config *viper.Viper = viper.New()

func GetDatabaseConfigProd() DatabaseConfig {
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("..")

	// membaca config
	err := config.ReadInConfig()
	HandleErrorWithPanic(err)
	// if err != nil {
	// 	panic(err.Error())
	// }

	databaseConfig := DatabaseConfig{
		Type:     config.GetString("database.prod.type"),
		Name:     config.GetString("database.prod.name"),
		Host:     config.GetString("database.prod.host"),
		Port:     strconv.Itoa(config.GetInt("database.prod.port")),
		Username: config.GetString("database.prod.username"),
		Password: config.GetString("database.prod.password"),
		Pool: DatabasePool{
			Idle: config.GetInt("database.prod.pool.idle"),
			Max: config.GetInt("database.prod.pool.max"),
			Lifetime: config.GetInt("database.prod.pool.lifetime"),
		},
	}
	return databaseConfig
}
func GetDatabaseConfigTest() DatabaseConfig {
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("..")

	// membaca config
	err := config.ReadInConfig()
	HandleErrorWithPanic(err)
	// if err != nil {
	// 	panic(err.Error())
	// }

	databaseConfig := DatabaseConfig{
		Type:     config.GetString("database.test.type"),
		Name:     config.GetString("database.test.name"),
		Host:     config.GetString("database.test.host"),
		Port:     strconv.Itoa(config.GetInt("database.test.port")),
		Username: config.GetString("database.test.username"),
		Password: config.GetString("database.test.password"),
		Pool: DatabasePool{
			Idle: config.GetInt("database.test.pool.idle"),
			Max: config.GetInt("database.test.pool.max"),
			Lifetime: config.GetInt("database.test.pool.lifetime"),
		},
	}
	return databaseConfig
}
