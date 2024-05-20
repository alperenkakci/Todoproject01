package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	UserRole     = "user"
	AdminRole    = "admin"
	DefaultUser1 = "user1"
	DefaultUser2 = "user2"
)

// AuthzMiddleware is the authorization middleware that checks user roles.
func AuthzMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check user roles
		if !checkRoles(claims.Username, roles) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func checkRoles(username string, roles []string) bool {
	for _, role := range roles {
		if role == AdminRole && (username == DefaultUser1 || username == DefaultUser2) {
			return true
		} else if role == UserRole && username != "" {
			return true
		}
	}
	return false
}
