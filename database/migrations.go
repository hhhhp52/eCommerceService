package database

//GORM migrations

import (
	"eCommerceService/src/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Migrate() {
	dsn := "root:123456@tcp(docker.for.mac.localhost:3306)/eCommerceService?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.UserToken{}, &models.Product{}, &models.RecommendationProduct{}, &models.RecommendationProductMapping{})
}
