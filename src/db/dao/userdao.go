package dao

import (
	"eCommerceService/src/models"
	"fmt"
	"github.com/jinzhu/gorm"
)

func SaveUser(db *gorm.DB, user models.User) error {
	// Insert user into the database
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("error inserting user: %v", err)
	}

	return nil
}

func UpdateUserVerified(db *gorm.DB, userID int, verified bool) error {
	// Update the verified field of the user
	if err := db.Model(&models.User{}).Where("id = ?", userID).Update("verified", verified).Error; err != nil {
		return fmt.Errorf("error updating user verified status: %v", err)
	}
	return nil
}

func CreateOrUpdateUserToken(db *gorm.DB, userToken models.UserToken) error {
	// Try to find the existing token
	var existingToken models.UserToken
	if err := db.Where("user_id = ?", userToken.UserID).First(&existingToken).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// If not found, create a new token
			if err := db.Create(&userToken).Error; err != nil {
				return fmt.Errorf("error inserting access token: %v", err)
			}
		} else {
			return fmt.Errorf("error finding access token: %v", err)
		}
	} else {
		// If found, update the existing token
		existingToken.AccessToken = userToken.AccessToken
		if err := db.Save(&existingToken).Error; err != nil {
			return fmt.Errorf("error updating access token: %v", err)
		}
	}

	return nil
}

func GetUserByToken(db *gorm.DB, token string) (models.User, error) {
	var user models.User
	if err := db.Table("user").Joins("JOIN user_token ON user.id = user_token.user_id").Where("user_token.access_token = ?", token).First(&user).Error; err != nil {
		return user, fmt.Errorf("error finding user by token: %v", err)
	}

	return user, nil
}

func GetUserByID(db *gorm.DB, id int) (*models.User, error) {
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, fmt.Errorf("error finding user by ID: %v", err)
	}
	return &user, nil
}
