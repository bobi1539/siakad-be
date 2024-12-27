package app

import (
	"database/sql"
	"fmt"
	"os"
	"siakad/constant"
	"siakad/helper"
	"time"

	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	godotenv.Load()
	driverName := os.Getenv(constant.DB_DRIVER)

	db, err := sql.Open(driverName, getDataSourceName())
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func getDataSourceName() string {
	dbUsername := os.Getenv(constant.DB_USERNAME)
	dbPassword := os.Getenv(constant.DB_PASSWORD)
	dbHost := os.Getenv(constant.DB_HOST)
	dbName := os.Getenv(constant.DB_NAME)

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbName)
}
