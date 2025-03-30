# app

A Gin web application with a clean architecture structure.

## Project Structure

```
.
├── controllers/     # HTTP request handlers
├── routes/         # Route definitions
├── services/       # Business logic
├── models/         # Data models
├── main.go         # Application entry point
├── go.mod          # Go module file
└── .gitignore      # Git ignore file
```

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

## API Endpoints

- `GET /`: Welcome message
- `GET /health`: Health check endpoint

## Development

The project follows a clean architecture pattern:
- Controllers handle HTTP requests
- Services contain business logic
- Models define data structures
- Routes define API endpoints

## License

MIT 