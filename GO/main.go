// main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/your-username/your-repo/auth"
	"github.com/your-username/your-repo/countries"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/auth", authenticateHandler).Methods("POST")
	r.HandleFunc("/country/{name}", tokenRequiredMiddleware(http.HandlerFunc(countryInfoHandler))).Methods("GET")
	r.HandleFunc("/countries", tokenRequiredMiddleware(http.HandlerFunc(countriesHandler))).Methods("GET")

	http.Handle("/", r)

	fmt.Printf("Server is running on http://localhost:%d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

// Authenticate handler calling auth package function
func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	// ... existing code

	token, err := auth.GenerateToken(credentials.Username, credentials.Password)
	if err != nil {
		// handle error
	}

	// ... existing code
}

// Country info handler and countries handler calling countries package functions
func countryInfoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	countryName := params["name"]

	// Call the relevant function from the 'countries' package
	countries.FetchCountryInfo(w, r, countryName)
}

func countriesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	population := r.FormValue("population")
	area := r.FormValue("area")
	language := r.FormValue("language")
	sort := r.FormValue("sort")
	page, _ := strconv.Atoi(r.FormValue("page"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	// Call the relevant function from the 'countries' package
	countries.FetchCountries(w, r, population, area, language, sort, page, limit)
}
