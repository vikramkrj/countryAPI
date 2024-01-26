package main

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

const (
	secretKey = "vik"
	username  = "vik"
	password  = "vik"
)

var (
	storedToken string
	router      *gin.Engine
)

// AuthRequest represents the structure of the JSON payload for authentication.
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	router = gin.Default()

	// Auth endpoint
	router.POST("/auth", authHandler)

	// Auth middleware for other endpoints
	authGroup := router.Group("/")
	authGroup.Use(tokenRequired)

	// Country details endpoint
	authGroup.GET("/country/:name", countryHandler)

	// Countries endpoint
	authGroup.GET("/countries", countriesHandler)

	port := 8080
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	router.Run(fmt.Sprintf(":%d", port))
}

// tokenRequired is a middleware to check the presence and validity of the JWT token.
func tokenRequired(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing!"})
		c.Abort()
		return
	}

	// Remove "Bearer " prefix from the token string
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Parse the token and check its validity
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token!"})
		c.Abort()
		return
	}

	// Continue with the next middleware or the handler
	c.Next()
}

// generateToken creates a new JWT token based on the provided username and password.
func generateToken(username, password string) (string, error) {
	if username == username && password == password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"password": password,
		})
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
	return "", fmt.Errorf("Invalid credentials!")
}

// authHandler handles the authentication endpoint.
func authHandler(c *gin.Context) {
	var authReq AuthRequest

	// Parse the JSON payload from the request
	if err := c.ShouldBindJSON(&authReq); err != nil {
		fmt.Println("Error parsing JSON payload:", err)
		fmt.Println("Received payload:", c.Request.Body)
		fmt.Printf("Received payload: %+v\n", authReq)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Generate a new token based on the provided credentials
	token, err := generateToken(authReq.Username, authReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Store the generated token
	storedToken = token
	c.JSON(http.StatusOK, gin.H{"storedToken": storedToken})
}

// countryHandler handles the country details endpoint.
func countryHandler(c *gin.Context) {
	countryName := c.Param("name")

	// Fetch detailed information about a specific country
	countryInfo, err := getCountryInfo(countryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch country information"})
		return
	}

	// Return the country information as JSON response
	c.JSON(http.StatusOK, countryInfo)
}

// countriesHandler handles the countries endpoint.
func countriesHandler(c *gin.Context) {
	// Retrieve a list of countries based on filters and sorting
	countries, err := getAllCountries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch countries"})
		return
	}

	// Return the list of countries as JSON response
	c.JSON(http.StatusOK, gin.H{"countries": countries})
}

// getCountryInfo fetches detailed information about a specific country.
func getCountryInfo(countryName string) (map[string]interface{}, error) {
	// Use the restcountries.com API or any other source to fetch detailed information about a specific country
	resp, err := resty.New().R().
		Get("https://restcountries.com/v3.1/name/" + countryName)
	if err != nil {
		return nil, err
	}

	// Extract the raw response body
	body := resp.Body()

	// Unmarshal the response body into a map
	var countryInfo map[string]interface{}
	if err := json.Unmarshal(body, &countryInfo); err != nil {
		return nil, err
	}

	return countryInfo, nil
}

// getAllCountries retrieves a list of all countries.
func getAllCountries() ([]map[string]interface{}, error) {
	// Use the restcountries.com API or any other source to fetch a list of countries
	resp, err := resty.New().R().
		Get("https://restcountries.com/v2/all")
	if err != nil {
		return nil, err
	}

	// Extract the raw response body
	body := resp.Body()

	// Unmarshal the response body into a slice of maps
	var countries []map[string]interface{}
	if err := json.Unmarshal(body, &countries); err != nil {
		return nil, err
	}

	return countries, nil
}