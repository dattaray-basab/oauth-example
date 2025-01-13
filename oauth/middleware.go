// middleware.go contains the middleware that will be used to authenticate the user.

package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func validateToken(token string) bool {
//     // Validate the token (e.g., with Google's /tokeninfo endpoint or JWT libraries)
//     return true // Placeholder logic
// }

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}
