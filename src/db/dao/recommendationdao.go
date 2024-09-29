package dao

import (
	"eCommerceService/src/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetRecommendationByUser(db *gorm.DB, userID int) ([]models.Product, error) {
	var products []models.Product
	err := db.Raw(`
    	SELECT product.* 
    	FROM recommendation_product_mapping 
    	JOIN product ON recommendation_product_mapping.recommendation_product_id = product.id 
    	WHERE recommendation_product_mapping.user_id = ?`, userID).Scan(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func SaveProduct(db *gorm.DB, product models.Product) error {
	// Insert product into the database
	if err := db.Create(&product).Error; err != nil {
		return fmt.Errorf("error inserting product: %v", err)
	}

	return nil
}

func SaveRecommendationProduct(db *gorm.DB, recommendationProduct models.RecommendationProduct) error {
	// Insert recommendation product into the database
	if err := db.Create(&recommendationProduct).Error; err != nil {
		return fmt.Errorf("error inserting recommendation product: %v", err)
	}

	return nil
}

func SaveRecommendationProductMapping(db *gorm.DB, recommendationProductMapping models.RecommendationProductMapping) error {
	// Insert recommendation product mapping into the database
	if err := db.Create(&recommendationProductMapping).Error; err != nil {
		return fmt.Errorf("error inserting recommendation product mapping: %v", err)
	}

	return nil

}
