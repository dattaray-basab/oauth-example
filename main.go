package main

import (
    "log"
    "os"

    "github.com/dattaray-basab/oauth-example/oauth"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Get values from environment variables
    port := os.Getenv("PORT")
    clientID := os.Getenv("CLIENT_ID")
    clientSecret := os.Getenv("CLIENT_SECRET")
    redirectURL := os.Getenv("REDIRECT_URL")

    // Call the RunOauth function
    oauth.RunOauth(port, map[string]string{
        "clientID":     clientID,
        "clientSecret": clientSecret,
        "redirectURL":  redirectURL,
    })
}
