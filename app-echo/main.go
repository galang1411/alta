package main

import (
	"app-echo/controller"
	"app-echo/domain/model"
	"app-echo/service"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = "root:gromizk123@tcp(127.0.0.1:3306)/crud_go?charset=utf8mb4&parseTime=True&loc=Local"
	}
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}

func main() {

	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	if err := MigrateDB(db); err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app := echo.New()
	us := service.NewDBUserService(db)
	controller.InitRouter(app, us)
	app.Start(":" + port)
}
