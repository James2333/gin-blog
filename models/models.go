package models

import (
	"fmt"
	"gin-blog/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err error
		 dbName, user, password, host string
		 tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	//dbType = sec.Key("mysql").String()
	//dbName = sec.Key("blog").String()
	//user = sec.Key("root").String()
	//password = sec.Key("root").String()
	//host = sec.Key("172.31.230.85").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	dbName="blog"
	user="root"
	password="root"
	host="172.31.230.85"

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	//db,err := gorm.Open("mysql","root:root@(172.31.230.85)/test?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName;
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

