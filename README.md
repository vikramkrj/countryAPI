# Go REST API with JWT Authentication

This is a simple Go implementation of a RESTful API with JWT authentication, inspired by a Node.js Express application. The API includes endpoints for user authentication, fetching detailed information about a specific country by name, and retrieving a list of countries based on filters and sorting.

## Features

- User authentication using JWT
- Token validation middleware for secure endpoints
- Endpoint to fetch detailed country information by name
- Endpoint to retrieve a list of countries based on filters and sorting

## Dependencies

- `github.com/gorilla/mux`: A powerful URL router and dispatcher for Go.
- `github.com/dgrijalva/jwt-go`: JSON Web Token implementation for Go.

## Installation

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