package db

import (
	"fmt"
	"time"

	"github.com/nakamurakzz/go-gin-api/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *SQLHandler
var err error

func DBOpen() {
	dbConn, err = NewSQLHandler()
	if err != nil {
		panic(err)
	}
}

func DBClose() {
	sqlDB, _ := dbConn.DB.DB()
	sqlDB.Close()
}

func NewSQLHandler() (*SQLHandler, error) {
	cnf := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cnf.DB_USER, cnf.DB_PASS, cnf.DB_HOST, cnf.DB_PORT, cnf.DB_DATABASE)

	var db *gorm.DB
	var err error

	// データベースに接続できるまでリトライする
	maxRetries := 10
	for i := 1; i <= maxRetries; i++ {
		fmt.Printf("Connecting to database... (try %d/%d)", i, maxRetries)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		if i < maxRetries {
			waitTime := time.Duration(i*2) * time.Second
			fmt.Printf("Failed to connect to database. Retrying in %v...\n", waitTime)
			time.Sleep(waitTime)
		}
	}

	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(30)
	sqlHandler := new(SQLHandler)

	db.Logger.LogMode(4)
	sqlHandler.DB = db

	return sqlHandler, nil
}

func GetDBConn() *SQLHandler {
	return dbConn
}
