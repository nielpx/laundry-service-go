package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// Define the specific token you want to validate
const validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA5NjYzMDUsInVzZXJuYW1lIjoidXNlcjEyMyJ9.NCAgvrE3S7uaINwVNeVX5lstPvY1hP_UYDs21PXXhbs"

// RequireAuth middleware checks for a valid token in the Authorization header
func RequireAuth(c *gin.Context) {
    // Get the Authorization header
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
        c.Abort()
        return
    }

    // Check if the token is prefixed with "Bearer "
    token := strings.TrimPrefix(authHeader, "Bearer ")
    if token == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
        c.Abort()
        return
    }

    // Validate if the token matches the specified valid token
    if token != validToken {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
        return
    }

    // Token is valid, continue to the next handler
    c.Next()
}
