package models

type Product struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (Product) TableName() string {
	return "product"
}

type RecommendationProduct struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	ProductID   int     `gorm:"not null" json:"product_id"`
	Product     Product `gorm:"foreignKey:ProductID" json:"product"`
	Description string  `json:"description"`
}

func (RecommendationProduct) TableName() string {
	return "recommendation_product"
}

type RecommendationProductMapping struct {
	ID                      int `gorm:"primaryKey" json:"id"`
	UserID                  int `gorm:"not null" json:"user_id"`
	RecommendationProductID int `gorm:"not null" json:"recommendation_product_id"`
}

func (RecommendationProductMapping) TableName() string {
	return "recommendation_product_mapping"
}
