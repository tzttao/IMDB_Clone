package Database

import (
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

var (
	userName = getEnv("DB_USER", "root")
	password = getEnv("DB_PASSWORD", "1234")
	dbName   = getEnv("DB_NAME", "IMDB")
	host     = getEnv("DB_HOST", "localhost")
	port     = getEnv("DB_PORT", "3306")
	charset  = getEnv("DB_CHARSET", "utf8")
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// DB Define global variable of database
var DB *gorm.DB

func init() {
	log.Println(">>>> get database connection start <<<<")
	DB = GetDataBase()
}

func GetDataBase() *gorm.DB {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	
	dsn := strings.Join([]string{userName, ":", password, "@tcp(", host, ":", port, ")/", dbName, "?charset=", charset, "&parseTime=True&loc=Local"}, "")
	
	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil
	}
	
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetMaxIdleConns(10)
	DB.SingularTable(true)
	
	debugMode := getEnv("DB_DEBUG", "false")
	DB.LogMode(debugMode == "true")
	
	log.Println("Database connection established successfully")
	return DB
}
