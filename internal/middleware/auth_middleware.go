package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

var secretKey = []byte(os.Getenv("SECRET"))

func RequireAuth(c *gin.Context) {
    // Get the cookie
    tokenString, err := c.Cookie("Authorization")
    if err != nil {
        logrus.Error("Cookie doesn't exist")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie not found"})
        c.Abort()
        return
    }

    // validasi token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return secretKey, nil
    })

    if err != nil || !token.Valid {
        logrus.Error("Invalid or expired token claims")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        c.Abort()
        return
    }

    // Ambil klaim
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        logrus.Error("Invalid token claims")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        c.Abort()
        return
    }

    username := claims["username"].(string)

    // Set user 
    c.Set("username", username)
    c.Next()
}
