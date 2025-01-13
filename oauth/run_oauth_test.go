package oauth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRunOauth(t *testing.T) {
	// Prepare test configurations
	port := ":8080"
	config := map[string]string{
		"clientID":     "test-client-id",
		"clientSecret": "test-client-secret",
		"redirectURL":  "http://localhost:8080/callback",
	}

	// Run the OAuth server in a goroutine
	go func() {
		RunOauth(port, config)
	}()

	// Create test cases for routes
	tests := []struct {
		route       string
		expectedCode int
	}{
		{route: "/", expectedCode: http.StatusOK},
		{route: "/login", expectedCode: http.StatusTemporaryRedirect},
		{route: "/protected/dashboard", expectedCode: http.StatusUnauthorized}, // AuthMiddleware will block it
	}

	// Allow the server some time to start (in production, you could handle this more robustly)
	serverURL := "http://localhost" + port
	for _, test := range tests {
		t.Run(test.route, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, serverURL+test.route, nil)
			w := httptest.NewRecorder()

			// Simulate the request
			http.DefaultServeMux.ServeHTTP(w, req)

			// Check the response status code
			if w.Code != test.expectedCode {
				t.Errorf("Route %s: expected status %d, got %d", test.route, test.expectedCode, w.Code)
			}
		})
	}
}
