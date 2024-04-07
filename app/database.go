package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB = ConnectionDatabase()

func ConnectionDatabase() *gorm.DB {
	dialect := mysql.Open("root:@tcp(127.0.0.1:3306)/go-simple-pos?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db
}
// migrate --database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up

