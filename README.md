# Basic HTTP Server

A simple HTTP server implementation using Go's `net/http` package. This project demonstrates how to create a basic web server with multiple endpoints that can serve both HTML and JSON responses.

## Features

- Simple HTTP server using Go's standard library
- Multiple endpoint handlers
- Serving static content
- JSON response capability

## Requirements

- Go 1.24 or higher

## Installation

```bash
# Clone the repository
git clone [repository-url]
cd [repository-name]
```

## Usage

```bash
# Run the server
go run main.go
```

Alternatively, you can use the provided Makefile:

```bash
make run
```

The server will start on port 8888. You can access the following endpoints:

- `http://localhost:8888/` - Home page
- `http://localhost:8888/about` - About page

## Project Structure

```
.
├── main.go        # Main server implementation
├── Makefile       # Build automation
├── README.md      # Project documentation
└── hint.md        # Additional hints and tips
```

## License

This project is licensed under the MIT License.
