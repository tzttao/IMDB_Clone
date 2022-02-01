package Database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

//数据库的基础信息
const (
	userName = "root"
	password = "GuoZiHan1998"
	dbName   = "mydb"
	host     = "localhost"
	port     = "3306"
	charset  = "utf8"
)

// 定义数据库全局变量
var Db *gorm.DB

func init() {
	log.Println(">>>> get database connection start <<<<")
	Db = GetDataBase()
}

func GetDataBase() *gorm.DB {
	// 给表名添加前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tp_" + defaultTableName
	}

	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=", charset}, "")
	//path := "root:GuoZiHan1998@tcp(localhost:3306)?charset=Utf8&parseTime=true&loc=Local"
	fmt.Println(path)
	//// 打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, err := gorm.Open("mysql", path)

	//DB, err := gorm.Open("mysql", "root:GuoZiHan1998@tcp(127.0.0.1:3306)/test")

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
