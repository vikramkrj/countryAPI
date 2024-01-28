# Go Countries Information Server

## Overview

This Go package provides a simple HTTP server with authentication and endpoints to retrieve information about countries. The server is built using the Gin framework and uses JWT (JSON Web Token) for authentication. It also interacts with the [restcountries.com](https://restcountries.com) API to fetch detailed information about countries.

## Installation

To use this package, you need to have Go installed on your system. Follow these steps:

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/go-resty/resty/v2


		
1	Import the package in your Go code:

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

Use the provided functionalities in your application.

2.	Usage
	
	2.1	Authentication
	The authentication endpoint /auth allows clients to obtain a JWT token by providing a valid username and password in the request body.

Example Request:

	bash
	curl -X POST -H "Content-Type: application/json" -d '{"username": "your_username", "password": "your_password"}' http://localhost:8080/auth

Example Response:

	json
	{
	  "storedToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InlvdXJfdXNlcm5hbWUiLCJwYXNzd29yZCI6InlvdXJfcGFzc3dvcmQifQ.NPT6dFsjbLvWz8s1PskZseXMU0tV2S8aSS6hovflbZc"
	}
	
	2.2	Country Information
	Country Details
	The endpoint /country/:name provides detailed information about a specific country.

Example Request:

	bash
	curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/country/germany

Example Response:

	json
	[
	  {
		"name": {
		  "common": "Germany",
		  "official": "Federal Republic of Germany",
		  "native": [
			"Deutschland"
		  ]
		},
		"tld": [
		  ".de"
		],
		"cca2": "DE",
		"ccn3": "276",
		"cca3": "DEU",
		"cioc": "GER",
		"independent": true,
		"status": "officially-assigned",
		"un_member": true,
		"currencies": {
		  "EUR": {
			"name": "Euro",
			"symbol": "€"
		  }
		},
		"idd": {
		  "root": "+4",
		  "suffixes": [
			"9"
		  ]
		},
		"capital": [
		  "Berlin"
		],
		"alt_spellings": [
		  "DE",
		  "Federal Republic of Germany",
		  "Bundesrepublik Deutschland"
		],
		"region": "Europe",
		"subregion": "Western Europe",
		"languages": {
		  "deu": "German"
		},
		"translations": {
		  "cym": "Yr Almaen",
		  "deu": "Deutschland",
		  "fra": "Allemagne",
		  "hrv": "Njemačka",
		  "ita": "Germania",
		  "jpn": "ドイツ",
		  "nld": "Duitsland",
		  "rus": "Германия",
		  "spa": "Alemania"
		},
		"latlng": [
		  51.0,
		  9.0
		],
		"demonym": "German",
		"landlocked": false,
		"borders": [
		  "AUT",
		  "BEL",
		  "CZE",
		  "DNK",
		  "FRA",
		  "LUX",
		  "NLD",
		  "POL",
		  "CHE"
		],
		"area": 357114.0,
		"flag": "🇩🇪",
		"flags": [
		  "https://restcountries.com/v3/flags/iso/flat/48/deu.png"
		]
	  }
	]
	
	2.3	List of Countries
The endpoint /countries provides a paginated list of countries based on optional filter parameters, sorting, and pagination.

Example Request:

	bash
	curl "Authorization: Bearer $TOKEN" http://localhost:8080/countries?population=100000000&area=500000&language=english&sort=name&order=asc&page=1&pageSize=10
Example Response:

	json
	{
	  "countries": [
		{
		  "name": {
			"common": "Canada",
			"official": "Canada"
		  },
		  "tld": [
			".ca"
		  ],
		  "cca2": "CA",
		  "ccn3": "124",
		  "cca3": "CAN",
		  "cioc": "CAN",
		  "independent": true,
		  "status": "officially-assigned",
		  "un_member": true,
		  "currencies": {
			"CAD": {
			  "name": "Canadian dollar",
			  "symbol": "$"
			}
		  },
		  "idd": {
			"root": "+1",
			"suffixes": [
			  "2"
			]
		  },
		  "capital": [
			"Ottawa"
		  ],
		  "alt_spellings": [
			"CA"
		  ],
		  "region": "Americas",
		  "subregion": "Northern America",
		  "languages": {
			"eng": "English",
			"fra": "French"
		  },
		  "translations": {
			"cym": "Canada",
			"deu": "Kanada",
			"fra": "Canada",
			"hrv": "Kanada",
			"ita": "Canada",
			"jpn": "カナダ",
			"nld": "Canada",
			"rus": "Канада",
			"spa": "Canadá"
		  },
		  "latlng": [
			60.0,
			-95.0
		  ],
		  "demonym": "Canadian",
		  "landlocked": false,
		  "borders": [
			"USA"
		  ],
		  "area": 9976140.0,
		  "flag": "🇨🇦",
		  "flags": [
			"https://restcountries.com/v3/flags/iso/flat/48/can.png"
		  ]
		},
		// ... additional countries ...
	  ]
	}

3.	Middleware
The package includes a middleware (tokenRequired) to check the presence and validity of the JWT token for protected endpoints. Ensure that the Authorization header contains a valid JWT token with the "Bearer" prefix.

4.	License
This package is licensed under the Prashant Advait Foundation. Feel free to modify and distribute