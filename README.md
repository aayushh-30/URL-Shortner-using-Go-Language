# URL Shortener in Go

This project is a simple URL shortener built using Go (Golang). It takes a long URL as input and returns a shortened URL, which can be used to redirect to the original URL.

## Features
- Generate shortened URLs using an MD5 hash.
- Store shortened URLs in an in-memory map.
- Redirect shortened URLs to their original URLs.
- Simple and fast API with JSON responses.
- Runs on `localhost:8000` by default.

## Endpoints

### 1. Health Check
**Endpoint:** `GET /`
**Description:** Check if the server is running.
**Response:**
```
Server is Running
```

### 2. Create Shortened URL
**Endpoint:** `POST /shorten`
**Request Body:**
```json
{
  "url": "https://example.com/long-url"
}
```
**Response:**
```json
{
  "short_url": "a1b2c3d4"
}
```

### 3. Redirect to Original URL
**Endpoint:** `GET /redirect/{short_url}`
**Description:** Redirects to the original URL associated with the given shortened URL.
**Response:**
- **302 Found:** Redirects to the original URL.
- **404 Not Found:** URL not found.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/url-shortener-go.git
   cd url-shortener-go
   ```
2. Run the server:
   ```bash
   go run main.go
   ```
3. The server will start at `http://localhost:8000`

## Testing
Use `curl` or any API testing tool like Postman to test the endpoints.
Example using `curl`:
```bash
curl -X POST http://localhost:8000/shorten -d '{"url": "https://example.com"}'
```

## License
This project is licensed under the MIT License.

