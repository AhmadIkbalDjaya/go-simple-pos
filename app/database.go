package app

import (
	"fmt"

	"github.com/AhmadIkbalDjaya/go-simple-pos/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB = ConnectionDatabase()

func ConnectionDatabase() *gorm.DB {
	
	// fmt.Println(util.Config.MysqlUsername)
	// fmt.Println(util.Config.MysqlPassword)
	// fmt.Println(util.Config.MysqlHost)
	// fmt.Println(util.Config.MysqlPort)
	// fmt.Println(util.Config.MysqlDatabase)
	// fmt.Printf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", util.Config.MysqlUsername, util.Config.MysqlPassword, util.Config.MysqlHost, util.Config.MysqlPort, util.Config.MysqlDatabase)
	dialect := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", util.Config.MysqlUsername, util.Config.MysqlPassword, util.Config.MysqlHost, util.Config.MysqlPort, util.Config.MysqlDatabase))
	// dialect := mysql.Open("root:@tcp(127.0.0.1:3306)/go_simple_pos?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db
}

func SetUpTestDatabase() *gorm.DB {
	dialect := mysql.Open("root:@tcp(127.0.0.1:3306)/go_simple_pos_test?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}
// migrate --database "mysql://root@tcp(localhost:3306)/go_simple_pos" -path db/migrations up

