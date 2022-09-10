package database

import (
	"fmt"
	"os"
	"path"

	"github.com/jinzhu/gorm"
	// import _ "github.com/jinzhu/gorm/dialects/mssql"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func OpenDbConnection() *gorm.DB {
	dialect := os.Getenv("DB_DIALECT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	var db *gorm.DB
	var err error
	if dialect == "sqlite3" {
		db, err = gorm.Open("sqlite3", path.Join(".", "app.db"))
	} else {
		databaseUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, username, password, dbName)
		db, err = gorm.Open(dialect, databaseUrl)
	}

	if err != nil {
		fmt.Print("db err: ", err)
		os.Exit(-1)
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}

func RemoveDb(db *gorm.DB) error {
	db.Close()
	err := os.Remove(path.Join(".", "app.db"))
	return err
}

func GetDb() *gorm.DB {
	return DB
}
