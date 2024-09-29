package seeder

import (
	"eCommerceService/src/db"
	"eCommerceService/src/db/dao"
	"eCommerceService/src/models"
	"fmt"
	"log"
)

func Seeder() {
	tx := db.DB()

	user, _ := dao.GetUserByID(tx, 1)
	if user != nil {
		fmt.Println("Database already seeded")
		return
	}

	// Seed Users
	users := []models.User{
		{Email: "user1@example.com", Password: "password1", Verified: true},
		{Email: "user2@example.com", Password: "password2", Verified: true},
		{Email: "user3@example.com", Password: "password3", Verified: true},
	}

	for _, user := range users {
		err := dao.SaveUser(tx, user)
		if err != nil {
			log.Fatalf("Error seeding users: %v", err)
		}
	}

	// Seed Products
	products := []models.Product{
		{Name: "Product1", Description: "Description1", Price: 10.0},
		{Name: "Product2", Description: "Description2", Price: 20.0},
		{Name: "Product3", Description: "Description3", Price: 30.0},
		{Name: "Product4", Description: "Description4", Price: 40.0}, // New product not in recommendation
	}

	for _, product := range products {
		err := dao.SaveProduct(tx, product)
		if err != nil {
			log.Fatalf("Error seeding products: %v", err)
		}
	}

	// Seed Recommendation Products
	recommendationProducts := []models.RecommendationProduct{
		{ProductID: 1, Description: "Recommendation for Product1"},
		{ProductID: 2, Description: "Recommendation for Product2"},
		{ProductID: 3, Description: "Recommendation for Product3"},
	}

	for _, recommendationProduct := range recommendationProducts {
		err := dao.SaveRecommendationProduct(tx, recommendationProduct)
		if err != nil {
			log.Fatalf("Error seeding recommendation products: %v", err)
		}
	}

	// Seed Recommendation Product Mappings
	recommendationProductMappings := []models.RecommendationProductMapping{
		{UserID: 1, RecommendationProductID: 1},
		{UserID: 1, RecommendationProductID: 2},
		{UserID: 2, RecommendationProductID: 2},
		{UserID: 2, RecommendationProductID: 3},
	}

	for _, recProductMapping := range recommendationProductMappings {
		err := dao.SaveRecommendationProductMapping(tx, recProductMapping)
		if err != nil {
			log.Fatalf("Error seeding recommendation product mappings: %v", err)
		}
	}

	fmt.Println("Seeding completed successfully")
}
