package database

import (
	"context"
)

func InitDetailedDatabase(ctx context.Context, filepath string) {
	WriteToMysql(ctx, filepath)
}

func InitOtherTables() {
	db := ContentMysql()
	// 只需要在程序开始时进行一次 AutoMigrate
	db.AutoMigrate(&Ec2Price{})

}
