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
	_ "golang-gorm-gin/cmd/app/docs"
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

// SignUp godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param request body models.User true "User registration data"
// @Success 200 {object} map[string]interface{} "Response when user created successfully"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 409 {object} map[string]interface{} "Conflict - Username already taken"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /signup [post]
func SignUp(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		logrus.Error("Invalid request data")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid request data"})
		return
	}

	var existingUser models.User
	if err := database.DB.Where("username = ?", newUser.Username).First(&existingUser).Error; err == nil {
		logrus.Error("Username alredy taken")
		c.JSON(http.StatusConflict, map[string]interface{}{"error": "Username already taken"})
		return
	}

	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		logrus.Error("Couldn't hash password")
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Could not hash password"})
		return
	}

	newUser.Password = hashedPassword
	if err := database.DB.Create(&newUser).Error; err != nil {
		logrus.Error("Couldn't create password")
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Could not create user"})
		return
	}
	logrus.Info("User created succesfully")
	c.JSON(http.StatusOK, map[string]interface{}{"message": "User created successfully"})
}

// LogIn godoc
// @Summary Log in a user
// @Description Authenticate a user and return a JWT token
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param request body models.User true "User login data"
// @Success 200 {object} map[string]interface{} "Response with token"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized - Invalid credentials"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /login [post]
func LogIn(c *gin.Context) {
	var credentials models.User
	if err := c.BindJSON(&credentials); err != nil {
		logrus.Error("Invalid request data")
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid request data"})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		logrus.Error("Invalid username or password")
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid username or password"})
		return
	}

	if !checkPassword(user.Password, credentials.Password) {
		logrus.Error("Invalid username or password")
		c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid username or password"})
		return
	}
	token, err := generateJWT(user.Username)
	if err != nil {
		logrus.Error("Couldn't generate JWT")
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Could not generate JWT"})
		return
	}
	logrus.Info("You are authorized")
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600, "", "", false, true)
	c.JSON(http.StatusOK, map[string]interface{}{"message": "Login successful", "token": token})
}
