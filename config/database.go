package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
	"fmt"
    "os"
)

func init(){
	err:=godotenv.Load()

	if err!=nil{
		fmt.Println("Error loading .env file")
	}
}

var DB *gorm.DB

func ConnectDatabase() {
    dsn := os.Getenv("DB_URL") // Example: "user:password@tcp(127.0.0.1:3306)/userdata?charset=utf8mb4&parseTime=True&loc=Local"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    DB = database

}
