package bootstrap

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type dbClient struct {
	KL *gorm.DB
	KC *gorm.DB
	KW *gorm.DB
}

func newDBClient(conf *sqlConf) *gorm.DB {
	dsn := fmt.Sprint("sqlserver://", conf.Username, ":", conf.Password, "@", conf.Host, ":", conf.Port, "?database=", conf.Database, "&parseTime=True")
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return db
}
