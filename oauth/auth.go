// auth.go contains the handlers for the OAuth server.

package oauth

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var GoogleOAuthConfig *oauth2.Config // Shared configuration

var state = "randomstate" // Replace with a securely generated random string

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to OAuth Example!"})
}

func Login(c *gin.Context) {
	authURL := GoogleOAuthConfig.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

func Callback(c *gin.Context) {
    if c.Query("state") != state {
        log.Println("Invalid state received:", c.Query("state"))
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OAuth state"})
        return
    }

    code := c.Query("code")
    token, err := GoogleOAuthConfig.Exchange(context.Background(), code)
    if err != nil {
        log.Println("Token exchange failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token.AccessToken})
}

func Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to your dashboard!"})
}
