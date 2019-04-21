package services

import (
	"github.com/ivahaev/go-logger"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"sync"
)

var instantiatedDB *gorm.DB
var onceDB sync.Once

var host string

//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
//check db connect in main function

func GetInstanceDB() *gorm.DB {
	host = "host=localhost port=5432 user=culture_admin dbname=culture_city password=qweASD123"
	onceDB.Do(func() {
		var err error
		instantiatedDB, err = gorm.Open("postgres", host)
		if err != nil {
			logger.Crit("Database initialization", err)
			os.Exit(64)
		}
	})
	return instantiatedDB
}
