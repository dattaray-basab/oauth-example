// run_oauth.go contains the main function to run the OAuth server.

package oauth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func RunOauth(port string, config map[string]string) {
	// Initialize OAuth configuration
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     config["clientID"],
		ClientSecret: config["clientSecret"],
		RedirectURL:  config["redirectURL"],
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	// Initialize Gin router
	router := gin.Default()

	// Public routes
	router.GET("/", Home)
	router.GET("/login", Login)
	router.GET("/callback", Callback)

	// Protected routes
	protected := router.Group("/protected")
	protected.Use(AuthMiddleware())
	protected.GET("/dashboard", Dashboard)

	// Start the server
	router.Run(port)
}
