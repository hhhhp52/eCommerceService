package models

import "time"

type User struct {
	ID              int       `gorm:"primaryKey" json:"id"`
	Email           string    `gorm:"unique;not null" json:"email"`
	Password        string    `gorm:"not null" json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
	Verified        bool      `gorm:"default:false" json:"verified"`
	VerifiedCode    string    `json:"verified_code"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}

type UserToken struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	UserID      int    `gorm:"not null" json:"user_id"`
	AccessToken string `gorm:"not null" json:"access_token"`
}

func (UserToken) TableName() string {
	return "user_token"
}
