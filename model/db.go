package model

import (
	"fmt"
	"goBlog/utlis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utlis.DbUser,
		utlis.DbPassWord,
		utlis.DbHost,
		utlis.DbPort,
		utlis.DbName)
	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		fmt.Printf("Failed to connect to the database. Please check the parameters", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		print(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//Ensure that the database connection timeout period is smaller than the network timeout period
	sqlDB.SetConnMaxLifetime(time.Second * 10)
}
