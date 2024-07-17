package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

// JWTVerify is the middleware to verify JWT tokens
//
// This middleware checks for a valid JWT token in the Authorization header of the request.
// If the token is missing, invalid, or has an unexpected signing method, the request is
// rejected with a 401 Unauthorized status. If the token is valid, the request proceeds
// to the next handler with the user information added to the context.
//
// Parameters:
//   - next: The next handler to be executed if the token is valid.
//
// Returns:
//   - An http.Handler that performs the JWT verification.
func JWTVerify(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Bearer token is missing", http.StatusUnauthorized)
			return
		}

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// Make sure that the token method conforms to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add user information to the context (optional)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), "user", claims["sub"]))

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

/*
// validateTokenHandler handles the /validate-token endpoint
// @Summary Validate ID token and generate JWT
// @Description Validate the provided ID token from Microsoft and generate a new JWT.
// @Tags auth
// @Accept json
// @Produce json
// @Param idToken body string true "ID Token"
// @Success 200 {string} string "JWT"
// @Failure 400 {string} string "Invalid token"
// @Failure 500 {string} string "Internal server error"
// @Router /api/v1/validateToken [post]
// @Security BearerAuth
func ValidateTokenHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		IDToken string `json:"idToken"`
	}

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Log the received ID token for debugging
    fmt.Printf("Received ID token: %s\n", req.IDToken)

	// Validate the ID token and generate a new JWT
	jwtString, err := controllers.ValidateTokenAndGenerateJWT(req.IDToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
    w.Write([]byte("Token validated successfully"))

	// Return the JWT
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"jwt": jwtString})
}
*/
