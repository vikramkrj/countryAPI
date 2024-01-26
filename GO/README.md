
# Country API Service

This backend API service provides useful data about countries using the REST Countries API (https://restcountries.com).

## Requirements

- Go (Golang)
- [Gin](https://github.com/gin-gonic/gin)
- [JWT-Go](https://github.com/dgrijalva/jwt-go)
- [Resty](https://github.com/go-resty/resty/v2)

## Setup

1. Clone the repository:

```bash
git clone <repository-url>
cd <repository-directory>
Install dependencies:
bash
go mod tidy
Run the application:
bash
go run main.go
The server will start running at http://localhost:8080.

API Endpoints
1. Auth Endpoint
Generate a valid auth token based on user credentials (username/password).

Endpoint:

bash

POST /auth
Payload:

json
{
  "username": "your_username",
  "password": "your_password"
}
2. Country Details Endpoint
Fetch detailed information about a specific country by providing its name as a parameter.

Endpoint:

bash
GET /country/:name
3. Countries Endpoint
Retrieve a list of all countries' names based on filters (population/area/language) and sorting (asc/desc). Support for pagination.

Endpoint:

bash
GET /countries
Query Parameters:

population
area
language
sort (asc/desc)
page
limit
API Authorization
All the above API endpoints are protected by authentication. Include the generated token in the Authorization header with the Bearer prefix.

Testing APIs
You can use curl commands to test the APIs from the command line. Below are some examples:

1. Auth Endpoint
bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"your_username","password":"your_password"}' http://localhost:8080/auth
2. Country Details Endpoint
bash
curl -H "Authorization: Bearer <your_generated_token>" http://localhost:8080/country/:name
3. Countries Endpoint
bash
curl -H "Authorization: Bearer <your_generated_token>" http://localhost:8080/countries?population=1000000&sort=desc&page=1&limit=10
Ensure to replace placeholders like your_username, your_password, <your_generated_token>, etc., with actual values.

Additional Information
For error handling and informative error messages, refer to the API responses.
Ensure to inspect the actual token being sent and received to ensure it matches the expected JWT format.
vbnet

Customize this `README.md` according to your project structure and requirements. If you have additional information or specific instructions, feel free to include them in the documentation.


