# Go REST API with JWT Authentication and Country Information

This Go project is a simple implementation of a RESTful API that includes JWT authentication and functionality to retrieve information about countries. The code is organized into separate packages for better modularity and maintainability.

## Features

1. **JWT Authentication:**
   - User authentication using JSON Web Tokens (JWT).
   - Middleware for token validation on secure endpoints.

2. **Country Information:**
   - Endpoint to fetch detailed information about a specific country by name.
   - Endpoint to retrieve a list of countries based on filters and sorting.

## Project Structure

The project is organized into the following packages:

1. **main:**
   - Contains the main server code.
   - Imports and utilizes functionality from other packages.

2. **auth:**
   - Handles user authentication and JWT token generation.
   - Defines the `Claims` struct for JWT claims.
   - Implements middleware for token validation.

3. **countries:**
   - Contains functionality related to country information.
   - Defines the `CountryInfo` struct for country details.
   - Implements functions to fetch country information and retrieve a list of countries.

## Dependencies

- [github.com/gorilla/mux](https://github.com/gorilla/mux): A powerful URL router and dispatcher for Go.
- [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go): JSON Web Token implementation for Go.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/your-repo.git


1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/your-repo.git

Install dependencies:

	''bash
	go get -u github.com/gorilla/mux
	go get -u github.com/dgrijalva/jwt-go

Run the application:

	go run main.go

The server will start on http://localhost:3000.

API Endpoints
1. Authentication
POST /auth
Authenticate the user and obtain a JWT token.

Request Body:

{
  "username": "your_username",
  "password": "your_password"
}

Response:

{
  "storedToken": "your_generated_token"
}

2. Fetch Country Information
GET /country/{name}
Fetch detailed information about a specific country by name.

3. Retrieve List of Countries
GET /countries
Retrieve a list of countries based on filters and sorting.

Query Parameters:

population
area
language
sort (optional, default is 'asc')
page (optional, default is 1)
limit (optional, default is 10)
Response:

{
  "totalPages": 1,
  "currentPage": 1,
  "countries": []
}

License
This project is licensed under the Prashnat Advait License.


Replace `https://github.com/your-username/your-repo.git` with the actual URL