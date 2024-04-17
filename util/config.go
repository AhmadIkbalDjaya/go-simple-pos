package util

import "github.com/spf13/viper"

type EnvConfig struct {
	AppPort string `mapstructure:"APP_PORT"`
	MysqlUsername string `mapstructure:"MYSQL_USERNAME"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MysqlHost string `mapstructure:"MYSQL_HOST"`
	MysqlPort string `mapstructure:"MYSQL_POrT"`
	MysqlDatabase string `mapstructure:"MYSQL_DATABASE"`
}

var Config *EnvConfig = &EnvConfig{}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.ReadInConfig()

	viper.Unmarshal(Config)
}