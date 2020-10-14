package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Services struct {
	ID    string `gorm:"primaryKey"`
	Value bool
	Key   string
}

var DB *gorm.DB

func Connector() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Unable to find the .env file")
		log.Fatal(err)
	}
	DB_User := os.Getenv("DB_USER")
	DB_Password := os.Getenv("DB_PASSWORD")
	DB_Port := os.Getenv("DB_PORT")
	DB_Name := os.Getenv("DB_NAME")
	dsn := DB_User + ":" + DB_Password + "@" + "(" + DB_Port + ")" + "/" + DB_Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection succesfull", DB)
	}
	// Makes the table of structure Services
	err = DB.AutoMigrate(&Services{})
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("built the table successfully")

}
