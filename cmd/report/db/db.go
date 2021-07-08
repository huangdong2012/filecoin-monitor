package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"grandhelmsman/filecoin-monitor/utils"
)

var (
	db      *gorm.DB
)

func Init(dbUrl string) {
	gdb, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		PrepareStmt: true,
	})
	utils.Throw(err)

	db = gdb
	sdb, err := gdb.DB()
	utils.Throw(err)
	sdb.SetMaxIdleConns(30)
	sdb.SetMaxOpenConns(100)
}
