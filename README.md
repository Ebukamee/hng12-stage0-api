# HNG12 Stage 0 API

This is a simple REST API built using Go and the Gorilla Mux router. The API returns a JSON response containing an email, the current UTC datetime, and a GitHub repository link.

## Features
- Serves a JSON response at the root endpoint (`/`)
- Uses `gorilla/mux` for routing
- Retrieves the current UTC time

## Technologies Used
- **Go**: The programming language used for development
- **Gorilla Mux**: A powerful HTTP router and dispatcher for Go
- **JSON Encoding**: Used to format the response in JSON

## Installation & Setup
### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official website](https://go.dev/).

### Clone the Repository
```sh
git clone https://github.com/ebukamee/hng12-stage0-api.git
cd hng12-stage0-api
```

### Install Dependencies
```sh
go mod tidy
```

### Run the Application
```sh
go run main.go
```

The server will start on port `8000`. You can access it by visiting:
```
http://localhost:8000/
```

## API Endpoint
### `GET /`
#### Response:
```json
{
  "email": "chukwuebukanwokike2007@gmail.com",
  "current_datetime": "2025-02-06T12:00:00Z",
  "github_url": "https://github.com/ebukamee/hng12-stage0-api"
}
```

## Project Structure
```
.
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## License
This project is open-source and available under the [MIT License](LICENSE).

## Author
**Chukwuebuka Nwokike**

For inquiries, reach out via email: `chukwuebukanwokike2007@gmail.com`.

