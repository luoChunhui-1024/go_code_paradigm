package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB 全局的数据库连接池
var (
	DB *sql.DB
)

// InitDB 初始化数据库连接
func InitDB() {
	dsn := "username:password@tcp(127.0.0.1:3306)/dbname"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// 尝试连接数据库，验证 DSN 是否正确
	if err := DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}
