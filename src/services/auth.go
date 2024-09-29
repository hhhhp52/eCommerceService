package services

import (
	"eCommerceService/src/db"
	"eCommerceService/src/db/dao"
	"eCommerceService/src/models"
	"eCommerceService/src/utils/helpers"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Registering User"})
			return
		}
	}()
	var user models.User
	tx := db.DB()
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check email
	if !helpers.IsValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Check password
	if !helpers.IsValidPassword(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be 6-16 characters long, contain one uppercase, one lowercase, and one special character"})
		return
	}

	// Check password confirmation
	if user.Password != user.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	verifiedCode := helpers.GenerateVerificationCode()
	user.VerifiedCode = verifiedCode

	err = dao.SaveUser(tx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user"})
		return
	}

	// Send verification email
	// for testing, we will not send the email
	//go helpers.SendVerificationEmail(user.Email, verifiedCode)

	// Set the Verified to true for testing
	err = dao.UpdateUserVerified(tx, user.ID, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.Status(http.StatusCreated)
}

func Login(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Login user"})
			return
		}
	}()
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate email and password
	if !helpers.IsValidEmail(input.Email) || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	var user models.User
	tx := db.DB()

	// Use parameterized query to prevent SQL injection
	if err := tx.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check if the user is verified
	if !user.Verified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not verified"})
		return
	}

	// Generate access token
	token, err := helpers.GenerateAccessToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
		return
	}

	// Save the access token
	userToken := models.UserToken{UserID: user.ID, AccessToken: token}
	if err := dao.CreateOrUpdateUserToken(tx, userToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

func VerifyEmail(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user"})
			return
		}
	}()
	var input struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var user models.User
	tx := db.DB()

	// Use parameterized query to prevent SQL injection
	if err := tx.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or code"})
		return
	}

	// Check if the verification code matches
	if user.VerifiedCode != input.Code {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or code"})
		return
	}

	// Set the user as verified
	if err := dao.UpdateUserVerified(tx, user.ID, true); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
