package helpers

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"regexp"
	"time"
)

func IsValidEmail(email string) bool {
	// Simple regex for email validation
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func IsValidPassword(password string) bool {
	// Password must be 6-16 characters long, contain one uppercase, one lowercase, and one special character
	if len(password) < 6 || len(password) > 16 {
		return false
	}
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)
	return hasLower && hasUpper && hasSpecial
}

// CheckPasswordHash compares a plain text password with a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateAccessToken generates a random access token
func GenerateAccessToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func SendVerificationEmail(email, code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("text/plain", "Please verify your email using this code: "+code)

	d := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-email-password")

	if err := d.DialAndSend(m); err != nil {
		// Handle error
		fmt.Println("Failed to send email:", err)
	}
}

func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
