package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const (
	PORT       = 3000
	SECRET_KEY = "vik"
	USERNAME   = "vik"
	PASSWORD   = "vik"
)

var storedToken string

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func generateToken(username, password string) (string, error) {
	if username == USERNAME && password == PASSWORD {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
			Username: username,
			Password: password,
		})
		return token.SignedString([]byte(SECRET_KEY))
	}
	return "", errors.New("Invalid credentials!")
}

func tokenRequiredMiddleware(next http.Handler) http.Handler {
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

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if credentials.Username == USERNAME && credentials.Password == PASSWORD {
		storedToken, err = generateToken(credentials.Username, credentials.Password)
		if err != nil {
			http.Error(w, "Token generation failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]string{"storedToken": storedToken}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid credentials!", http.StatusUnauthorized)
	}
}

func countryInfoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	countryName := params["name"]

	response, err := http.Get("https://restcountries.com/v3.1/name/" + countryName)
	if err != nil {
		http.Error(w, "Failed to fetch country information!", http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&w)
}

func countriesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	population := r.FormValue("population")
	area := r.FormValue("area")
	language := r.FormValue("language")
	sort := r.FormValue("sort")
	page, _ := strconv.Atoi(r.FormValue("page"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	// Your existing logic for fetching and filtering countries goes here

	// For demonstration purposes, returning a dummy response
	response := map[string]interface{}{
		"totalPages":   1,
		"currentPage":  page,
		"countries":    []interface{}{},
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/auth", authenticateHandler).Methods("POST")
	r.HandleFunc("/country/{name}", tokenRequiredMiddleware(http.HandlerFunc(countryInfoHandler))).Methods("GET")
	r.HandleFunc("/countries", tokenRequiredMiddleware(http.HandlerFunc(countriesHandler))).Methods("GET")

	http.Handle("/", r)

	fmt.Printf("Server is running on http://localhost:%d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
