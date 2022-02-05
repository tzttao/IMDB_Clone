package Database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

//var MysqlDb *sql.DB
//var MysqlDbErr error
//
//const (
//	USER_NAME = "root"
//	PASS_WORD = "root123456"
//	HOST      = "localhost"
//	PORT      = "3306"
//	DATABASE  = "IMDB"
//	CHARSET   = "utf8"
//)

//// 初始化链接
//func init() {
//
//	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
//
//	// 打开连接失败
//	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
//	//defer MysqlDb.Close();
//	if MysqlDbErr != nil {
//		log.Println("dbDSN: " + dbDSN)
//		panic("数据源配置不正确: " + MysqlDbErr.Error())
//	}
//
//	// 最大连接数
//	MysqlDb.SetMaxOpenConns(100)
//	// 闲置连接数
//	MysqlDb.SetMaxIdleConns(20)
//	// 最大连接周期
//	MysqlDb.SetConnMaxLifetime(100 * time.Second)
//
//	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
//		panic("数据库链接失败: " + MysqlDbErr.Error())
//	}
//
//}

//数据库的基础信息
const (
	userName = "root"
	password = "root123456"
	dbName   = "IMDB"
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
		return defaultTableName
	}

	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=", charset}, "")
	fmt.Println(path)
	// 打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
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
