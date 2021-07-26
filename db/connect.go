package db

import (
	"fmt"
	"github.com/CoryKelly/Admin_App/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Database Gorm DataBase Pointer
var Database *gorm.DB

func Connect() {
	//Connection
	db, err := gorm.Open(mysql.Open("root:nikolai365@/go_admin"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected...")

	//Export DB value
	Database = db

	//Importing models
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{})
}
