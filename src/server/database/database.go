package database

import (
	"fmt"
	"todo-app/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func InitializeDatabase() *Database {
	dsn := "root:root@tcp(127.0.0.1:3306)/todo-app?charset=utf8mb4&parseTime=True&loc=Local"
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
