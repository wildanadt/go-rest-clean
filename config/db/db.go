package db

import (
	"fmt"

	"github.com/wildanadt/go-rest-clean/config/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var dbConfig = env.Config.PostgresConfig

func init() {

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	DB = db
}

func GetDBInstance() *gorm.DB {
	// fmt.Println(DB)
	return DB

}
