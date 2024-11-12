package handler

import (
	"golang-gorm-gin/internal/database"
	"golang-gorm-gin/internal/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error("Something went wrong while hashing password")
		return "", err
	}
	logrus.Info("Hashed successfully")
	return string(hash), nil
}

func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

var secretKey = []byte(os.Getenv("SECRET"))

func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Expires in 72 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}


func SignUp(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		logrus.Error("Invalid request data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Check kalau user ada
	var existingUser models.User
	if err := database.DB.Where("username = ?", newUser.Username).First(&existingUser).Error; err == nil {
		logrus.Error("Username alredy taken")
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Hash the password sebelum storing
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		logrus.Error("Couldn't hash password")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Store user in the database
	newUser.Password = hashedPassword
	if err := database.DB.Create(&newUser).Error; err != nil {
		logrus.Error("Couldn't create password")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	logrus.Info("User created succesfully")
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Login route
func LogIn(c *gin.Context) {
	var credentials models.User
	if err := c.BindJSON(&credentials); err != nil {
		logrus.Error("Invalid request data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Check kalau user ada
	var user models.User
	if err := database.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		logrus.Error("Invalid username or password")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Check password
	if !checkPassword(user.Password, credentials.Password) {
		logrus.Error("Invalid username or password")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	// Generate token
	token, err := generateJWT(user.Username)
	if err != nil {
		logrus.Error("Couldn't generate JWT")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate JWT"})
		return
	}
	logrus.Info("You are authorized")
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
