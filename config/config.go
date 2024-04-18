package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	AppPort 			string `mapstructure:"APP_PORT"`
	MysqlUsername string `mapstructure:"MYSQL_USERNAME"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MysqlHost 		string `mapstructure:"MYSQL_HOST"`
	MysqlPort 		string `mapstructure:"MYSQL_PORT"`
	MysqlDatabase string `mapstructure:"MYSQL_DATABASE"`
}

var Config *EnvConfig = LoadConfig()

func LoadConfig() (config *EnvConfig) {
	viper.SetDefault("APP_PORT", "3003")
	viper.SetDefault("MYSQL_USERNAME", "root")
	viper.SetDefault("MYSQL_PASSWORD", "")
	viper.SetDefault("MYSQL_HOST", "127.0.0.1")
	viper.SetDefault("MYSQL_PORT", "3306")
	viper.SetDefault("MYSQL_DATABASE", "go_simple_pos")

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.ReadInConfig()

	viper.Unmarshal(&config)
	return 
}