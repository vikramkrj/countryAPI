const express = require('express');
const bodyParser = require('body-parser');
const jwt = require('jsonwebtoken');
const axios = require('axios');

const app = express();
const PORT = 3000;
const SECRET_KEY = 'vik';

// Hardcoded username and password for simplicity
const USERNAME = 'vik';
const PASSWORD = 'vik';

let storedToken;
// Middleware for token authentication
const tokenRequired = (req, res, next) => {
  const token = req.headers.authorization;
	const token1 = generateToken(USERNAME, PASSWORD)
  if (!token) {
    return res.status(401).json({ error: 'Token is missing!' });
  }

  try {
    // In a real-world scenario, use a library like jsonwebtoken for secure token validation
    if (token !== storedToken) {
      throw new Error('Invalid token!');
    }
    next();
  } catch (error) {
    return res.status(401).json({error:  '${token} Invalid token!' });
  }
};

// Function to generate JWT token
const generateToken = (username, password) => {
  // In a real-world scenario, use a library like jsonwebtoken for secure token generation
  if (username === USERNAME && password === PASSWORD) {
    return jwt.sign({ username, password }, SECRET_KEY, { expiresIn: '1h' });
  } else {
    throw new Error('Invalid credentials!');
  }
};

// API endpoint for authentication
app.post('/auth', bodyParser.json(), (req, res) => {
  const { username, password } = req.body;

  try {
    if (username === USERNAME && password === PASSWORD) {
      storedToken = generateToken(username, password);
      res.json({ storedToken });
    } else {
      res.status(401).json({ error: 'Invalid credentials!' });
    }
  } catch (error) {
    res.status(401).json({ error: 'Invalid credentials!' });
  }
});

// API endpoint to fetch detailed information about a specific country by name
app.get('/country/:name', tokenRequired, async (req, res) => {
//	app.get('/country/:name', async (req, res) => {
  const countryName = req.params.name;

  try {
    const response = await axios.get(`https://restcountries.com/v3.1/name/${countryName}`);
    res.json(response.data);
  } catch (error) {
    res.status(500).json({ error: 'Failed to fetch country information!' });
  }
});

// API endpoint to retrieve a list of countries based on filters and sorting
//app.get('/countries', tokenRequired, async (req, res) => {
	app.get('/countries', async (req, res) => {
  const { population, area, language, sort, page, limit } = req.query;
  const queryParams = new URLSearchParams({
    population,
    area,
    language,
    sort,
    page,
    limit,
  });

  try {
    const response = await axios.get(`https://restcountries.com/v3.1/all?${queryParams}`);
	//https://restcountries.com/v3.1/independent?status=true
    res.json(response.data);
  } catch (error) {
    res.status(500).json({ error: 'Failed to fetch countries!' });
  }
});

// Start the server
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
