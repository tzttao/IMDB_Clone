package Database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

const (
	userName = "root"
	password = "root123456"
	dbName   = "IMDB"
	host     = "localhost"
	port     = "3306"
	charset  = "utf8"
)

// DB Define global variable of database
var DB *gorm.DB

func init() {
	log.Println(">>>> get database connection start <<<<")
	DB = GetDataBase()
}

func GetDataBase() *gorm.DB {
	// 给表名添加前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	// Build connection
	path := strings.Join([]string{userName, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=", charset}, "")
	//fmt.Println(path)
	// Open database. 前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, err := gorm.Open("mysql", path)
	if err != nil {
		log.Fatal("连接失败！")
		return nil
	}
	// 设置数据库最大连接数
	DB.DB().SetMaxOpenConns(500)
	// 设置数据库最大闲置数
	DB.DB().SetMaxIdleConns(20)
	// 全局禁用表名复数
	DB.SingularTable(true)
	// 调试模式，可以打印sql语句
	DB.LogMode(true)
	return DB
}
