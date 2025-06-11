# go-custom-http

A custom HTTP server and router implementation in Go.

## Overview

This project demonstrates a simple HTTP server using a custom HTTP library. It includes basic routing and handler registration for endpoints.

## Features

- Custom router creation
- Handler registration for specific paths
- Basic request and response structures
- Example endpoints:
  - `/hello`: Returns a success message
  - `/helloError`: Returns an error message

## Usage

Clone the repository and run:

```bash
go run main.go
```

The server will listen on port `8080`. You can test the endpoints with:

- `http://localhost:8080/hello`
- `http://localhost:8080/helloError`

## Project Structure

- `main.go`: Entry point, sets up the server and routes.
- `custom-http`: Contains the custom HTTP handling logic (imported from the `github.com/brysonurie/go-http/custom-http` package).

## Requirements

- Go 1.18 or newer

## License

MIT License (or specify your license here)
