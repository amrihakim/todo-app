package database

import (
	"fmt"
	"os"
	"todo-app/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func InitializeDatabase() *Database {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv(`MYSQL_USER`),
		os.Getenv(`MYSQL_PASSWORD`),
		os.Getenv(`MYSQL_HOST`),
		os.Getenv(`MYSQL_PORT`),
		os.Getenv(`MYSQL_DBNAME`))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open connection")
		panic(err)
	}

	db.AutoMigrate(
		&models.Activity{},
		&models.Todo{},
	)

	return &Database{db}
}
