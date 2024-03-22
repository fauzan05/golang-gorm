package helper

import (
	"strconv"

	"github.com/spf13/viper"
)

type DatabaseConfig struct{
	Type string
	Name string
	Host string
	Port string
	Username string
	Password string
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
		Type: config.GetString("database.prod.type"),
		Name: config.GetString("database.prod.name"),
		Host: config.GetString("database.prod.host"),
		Port: strconv.Itoa(config.GetInt("database.prod.port")),
		Username: config.GetString("database.prod.username"),
		Password: config.GetString("database.prod.password"),
	}
	return databaseConfig
}
func GetDatabaseConfigTest() DatabaseConfig {
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// membaca config
	err := config.ReadInConfig()
	HandleErrorWithPanic(err)
	// if err != nil {
	// 	panic(err.Error())
	// }
	
	databaseConfig := DatabaseConfig{
		Type: config.GetString("database.test.type"),
		Name: config.GetString("database.test.name"),
		Host: config.GetString("database.test.host"),
		Port: strconv.Itoa(config.GetInt("database.test.port")),
		Username: config.GetString("database.test.username"),
		Password: config.GetString("database.test.password"),
	}
	return databaseConfig
}