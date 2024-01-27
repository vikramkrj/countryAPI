package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"
	"sort"
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

// parseFilterParameters extracts and parses filter parameters from the request query.
func parseFilterParameters(c *gin.Context) (int, int, string, string, string, int, int) {
	populationStr := c.DefaultQuery("population","")
	areaStr := c.DefaultQuery("area","")

	var populationFilter, areaFilter int
	var err error

	if populationStr != "" {
		populationFilter, err = strconv.Atoi(populationStr)
		if err != nil {
			fmt.Println("Error parsing population:", err)
			// Handle the error as needed, maybe set a default value
			populationFilter = 0
		}
	}

	if areaStr != "" {
		areaFilter, err = strconv.Atoi(areaStr)
		if err != nil {
			fmt.Println("Error parsing area:", err)
			// Handle the error as needed, maybe set a default value
			areaFilter = 0
		}
	}

	languageFilter := c.DefaultQuery("language","")
	if languageFilter == "" {
		fmt.Println("Language filter is empty. No language filter will be applied.")
	}

	sortBy := c.DefaultQuery("sort","")
	sortOrder := c.DefaultQuery("order", "asc") // Default to ascending order
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		fmt.Println("Error parsing page:", err)
		// Handle the error as needed, maybe set a default value
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10")) // Default page size to 10
	if err != nil {
		fmt.Println("Error parsing pageSize:", err)
		// Handle the error as needed, maybe set a default value
		pageSize = 10
	}

	return populationFilter, areaFilter, languageFilter, sortBy, sortOrder, page, pageSize
}

// countriesHandler handles the countries endpoint.
func countriesHandler(c *gin.Context) {

	// Print entire request URL for debugging
	fmt.Println("Request URL:", c.Request.URL.String())

	// Retrieve filter parameters from the request query
	populationFilter, areaFilter, languageFilter, sortBy, sortOrder, page, pageSize := parseFilterParameters(c)

	// Print filter parameters for debugging
	fmt.Println("Filter Parameters:\nPopulation: %d\nArea: %d\nLanguage: %s\nSortBy: %s\nSortOrder: %s\nPage: %d\nPageSize: %d\n",
		populationFilter, areaFilter, languageFilter, sortBy, sortOrder, page, pageSize)

	// Retrieve a list of countries based on filters, sorting, and pagination
	countries, err := getAllCountries(populationFilter, areaFilter, languageFilter, sortBy, sortOrder, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch countries"})
		return
	}

	// Return the list of countries as JSON response
	c.JSON(http.StatusOK, gin.H{"countries": countries})
}



// getCountryInfo fetches detailed information about a specific country.
func getCountryInfo(countryName string) ([]map[string]interface{}, error) {
	// Use the restcountries.com API or any other source to fetch detailed information about a specific country
	resp, err := resty.New().R().
		Get("https://restcountries.com/v3.1/name/" + countryName)
	if err != nil {
		fmt.Println("Error fetching country information:", err)
		return nil, err
	}

	// Extract the raw response body
	body := resp.Body()

	// Print the raw response body for debugging
	fmt.Println("Raw response body:", string(body))

	// Unmarshal the response body into a map
	var countryInfo []map[string]interface{}
	if err := json.Unmarshal(body, &countryInfo); err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return nil, err
	}


	// Print the country information for debugging
	fmt.Println("Country Information:", countryInfo)
	return countryInfo, nil
}

// getAllCountries retrieves a list of all countries based on filters, sorting, and pagination.
func getAllCountries(populationFilter int, areaFilter int, languageFilter string, sortBy string, sortOrder string, page int, pageSize int) ([]map[string]interface{}, error) {
	// Use the restcountries.com API or any other source to fetch a list of all countries
	resp, err := resty.New().R().
		Get("https://restcountries.com/v3/all")
	if err != nil {
		log.Println("Error fetching countries:", err)
		return nil, err
	}

	// Extract the raw response body
	body := resp.Body()

	// Unmarshal the response body into a slice of maps
	var allCountries []map[string]interface{}
	if err := json.Unmarshal(body, &allCountries); err != nil {
		log.Println("Error unmarshalling response body:", err)
		return nil, err
	}

// Apply filters
var filteredCountries []map[string]interface{}
for _, country := range allCountries {
    // Apply population filter
if populationFilter > 0 {
   population, popFound := country["population"].(float64)
    if !popFound {
  //      fmt.Printf("No population field found for country: %v\n", country["name"])
        continue
    }

    if populationFilter > 0 && population > float64(populationFilter) {
   //     fmt.Printf("Skipped country %s due to population filter: %v\n", country["name"], country)
        continue
    }
}
 

    // Apply area filter
if areaFilter > 0  {
  // Apply area filter
area, areaFound := country["area"].(float64)
if !areaFound {
    fmt.Printf("No area field found for country: %v\n", country["name"])
    continue
}

if areaFilter > 0 && area > float64(areaFilter) {
    fmt.Printf("Skipped country %s due to area filter: %v\n", country["name"], country)
    continue
}

}

    // Apply language filter
if languageFilter != "" {
    languagesMap, langFound := country["languages"].(map[string]interface{})
    if !langFound {
        fmt.Printf("No languages field found for country: %v\n", country["name"])
        continue
    }

    languageFound := false
    for langCode, langName := range languagesMap {
        // Check if the language code and name match the filter
        if langCode == languageFilter || langName == languageFilter {
            languageFound = true
            break
        }
    }

    if !languageFound {
        fmt.Printf("Skipped country %s due to language filter: \n", country["name"])
        continue
    }
	if languageFound {
	fmt.Println("Country found %s :", country["name"])
	}
}
    // If all filters pass, add the country to the filtered list
    filteredCountries = append(filteredCountries, country)
}


// Apply sorting
sort.SliceStable(filteredCountries, func(i, j int) bool {
    // Ensure the fields exist
    fieldI, fieldJ := filteredCountries[i][sortBy], filteredCountries[j][sortBy]

    // Print debug information
    fmt.Printf("Sorting fields for countries:\n")
    fmt.Printf("Country %s: %v, Field %s: %v\n", filteredCountries[i]["name"], filteredCountries[i], sortBy, fieldI)
    fmt.Printf("Country %s: %v, Field %s: %v\n", filteredCountries[j]["name"], filteredCountries[j], sortBy, fieldJ)

    // Handle different types
    switch fieldI := fieldI.(type) {
    case string:
        switch fieldJ := fieldJ.(type) {
        case string:
            // Compare string values
            if sortOrder == "asc" {
                return fieldI < fieldJ
            } else {
                return fieldI > fieldJ
            }
        default:
            return false
        }
    case int, float64:
        // Handle numeric types
        valueI, okI := convertToFloat(fieldI)
        valueJ, okJ := convertToFloat(fieldJ)

        if !okI || !okJ {
            fmt.Printf("Error converting numeric fields for countries:\n")
            return false
        }

        if sortOrder == "asc" {
            return valueI < valueJ
        } else {
            return valueI > valueJ
        }
    default:
        // Unsupported type, consider handling other types as needed
        return false
    }
})


	// Apply pagination
startIndex := (page - 1) * pageSize
endIndex := startIndex + pageSize

if startIndex >= len(filteredCountries) {
    // If startIndex is beyond the length of the filteredCountries, set both indices to len(filteredCountries)
    startIndex = len(filteredCountries)
    endIndex = len(filteredCountries)
} else if endIndex > len(filteredCountries) {
    // If endIndex is beyond the length of the filteredCountries, set it to len(filteredCountries)
    endIndex = len(filteredCountries)
}


	//fmt.Printf("Filtered Countries: %v\n", filteredCountries)
fmt.Println("start index", startIndex)
fmt.Println("EndIndex", endIndex)
	return filteredCountries[startIndex:endIndex], nil
}


// Function to convert interface{} to float64
func convertToFloat(value interface{}) (float64, bool) {
    switch v := value.(type) {
    case int:
        return float64(v), true
    case float64:
        return v, true
    default:
        return 0, false
    }
}