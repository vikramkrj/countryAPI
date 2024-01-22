
# README  for the  Country Information API

This project is an Express.js API that provides information about countries. It includes authentication and endpoints to retrieve detailed information about a specific country, as well as a list of countries with filtering and sorting options.

## Setup

1. Make sure you have Node.js and npm installed on your system.
2. Clone the repository:

			`` bash 
			git clone <repository-url>
			cd <repository-directory>
 
3 .Install dependencies:
			`` bash
			npm install

4. Start the server:
			`` bash
			npm start
The server will be running on http://localhost:3000.

5.	Authentication
	To authenticate and obtain a token, make a POST request to /auth with the following curl command:

			`` bash
			curl -X POST -H "Content-Type: application/json" -d '{"username":"vik", "password":"vik"}' http://localhost:3000/auth
	This will return a JSON object containing the storedToken.

6.	Country Information
	Get Detailed Information for a Country
	To get detailed information about a specific country, make a GET request to /country/:name with a valid token:

			``bash
			curl -H "Authorization: <storedToken>" http://localhost:3000/country/<country-name>
			Replace <storedToken> with the token obtained during authentication and <country-name> with the desired country's name.

7.	Get List of Countries with Filters and Sorting
	To retrieve a list of countries based on filters and sorting, make a GET request to /countries with a valid token:

			``bash
			curl -H "Authorization: <storedToken>" "http://localhost:3000/countries?population=<population>&area=<area>&language=<language>&sort=<asc/desc>&page=<page>&limit=<limit>"
			Replace <storedToken> with the token obtained during authentication, and provide optional query parameters for filtering and sorting.

			<population>: Population filter
			<area>: Area filter
			<language>: Language filter
			<asc/desc>: Sorting order (asc/desc)
			<page>: Page number for pagination
			<limit>: Number of items per page
	Example
		Here's an example command to get a list of countries with English as the language, sorted by population in descending order:

			``bash
			curl -H "Authorization: <storedToken>" "http://localhost:3000/countries?language=English&sort=desc&page=1&limit=10"
			Feel free to adjust the parameters based on your requirements.

	Make sure to replace <repository-url> and <repository-directory> with the actual URL and directory of your repository, respectively. 
	This README.md file provides instructions for setting up, authenticating, and making requests to your Express.js API using curl commands.
