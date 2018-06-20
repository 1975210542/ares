package models

import (
	"fmt"

	"net/url"

	"onlineshop/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DB = gorm.DB

var (
	db      *DB
	migrate = []interface{}{}
)

func init() {
	db = CreateDB()
	RunMigrate()
}

func connectString() string {

	host := common.GetConfig("mysql::host")
	port := common.GetConfig("mysql::port")
	user := common.GetConfig("mysql::user")
	password := common.GetConfig("mysql::password")
	database := common.GetConfig("mysql::database")
	param := "?"
	loc := url.QueryEscape("Asia/Shanghai")
	fmt.Println(host, port, user, password, database)
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%sloc=%s&charset=utf8&parseTime=true",
		user, password, host, port, database, param, loc)
}

func CreateDB() *DB {
	db, e := gorm.Open("mysql", connectString())
	if e != nil {
		panic(e)
	}
	return db

}

func ORM() *DB {
	if db == nil || db.DB().Ping() != nil {
		db = CreateDB()
	}

	return db
}

func SetMigrate(table interface{}) {
	migrate = append(migrate, table)
	fmt.Println("databases111:", migrate)

}

func RunMigrate() {
	fmt.Println("start....")
	for _, v := range migrate {
		fmt.Println("vVVVV:", v)
		if !ORM().HasTable(v) {
			ORM().AutoMigrate(v)
		}

	}
}
