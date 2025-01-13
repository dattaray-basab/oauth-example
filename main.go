// main.go

package main

import (
	"os"

	"github.com/dattaray-basab/oauth-example/oauth"
)

func main() {
	// Call the RunOauth function with parameters

	oauth.RunOauth(":8080", map[string]string{
		"clientID":     os.Getenv("GOOGLE_CLIENT_ID"),
		"clientSecret": os.Getenv("GOOGLE_CLIENT_SECRET"),
		"redirectURL":  "http://localhost:8080/callback",
	})

}
