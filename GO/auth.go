// auth/auth.go

package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Claims represents the structure for JWT claims.
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// SECRET_KEY, USERNAME, and PASSWORD constants are already defined in the main.go file

// GenerateToken generates a JWT token for a given username and password.
func GenerateToken(username, password string) (string, error) {
	if username == USERNAME && password == PASSWORD {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
			Username: username,
			Password: password,
		})
		return token.SignedString([]byte(SECRET_KEY))
	}
	return "", errors.New("Invalid credentials!")
}

// TokenRequiredMiddleware is a middleware for token authentication.
func TokenRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")

		if token == "" {
			http.Error(w, "Token is missing!", http.StatusUnauthorized)
			return
		}

		claims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		})

		if err != nil || !parsedToken.Valid {
			http.Error(w, fmt.Sprintf("Invalid token! %s", err), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// AuthenticateHandler handles user authentication and generates a JWT token.
func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := GenerateToken(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Token generation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"storedToken": token}
	json.NewEncoder(w).Encode(response)
}
