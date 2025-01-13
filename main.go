package main

import "github.com/dattaray-basab/oauth-example/oauth"

func main() {
	// Call the RunOauth function with parameters
	oauth.RunOauth(":8080", map[string]string{
		"clientID":     "your-google-client-id",
		"clientSecret": "your-google-client-secret",
		"redirectURL":  "http://localhost:8080/callback",
	})
}
