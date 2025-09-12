package database

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"log"

	"analysis-aws-bill/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ContentMysql() *gorm.DB {
	//读取.ini里面的数据库配置

	//dsn := fmt.Sprintf("root:ql2252528@tcp(127.0.0.1:3306)/bill202502?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MysqlUser, config.MysqlPassword, config.MysqlHost, config.MysqlPort, config.MysqlDbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:            true, //打印sql
		SkipDefaultTransaction: true,
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get SQL DB instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(6532)
	sqlDB.SetMaxIdleConns(10)

	return DB
}
