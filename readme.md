# Student API

This project is a simple API service for handling student data, built in Go. The service includes basic setup for a web server with graceful shutdown, router configuration, and a student handler.

## Features
- **Configuration Loading:** Automatically loads configuration using `config.MustLoad()`.
- **HTTP Server Setup:** Sets up an HTTP server using Go's `net/http` package.
- **Routing:** Adds a route for handling student-related requests (currently set up for `POST /api/students`).
- **Graceful Shutdown:** Listens for OS interrupt signals (SIGINT, SIGTERM) to gracefully shut down the server, allowing for a clean exit and freeing up resources.

## Project Structure
- **config**: Loads and provides configuration settings.
- **internal/http/handlers/students**: Contains HTTP handlers, such as the `New()` handler for student creation.

## Prerequisites
- Go 1.16+ installed
- Required packages can be installed using `go mod tidy`.

## How to Run
1. Clone the repository:
   ```bash
   git clone https://github.com/BibhabenduMukherjee/student-api.git


## Key Code Components
### Server Initialization
The server is initialized with an HTTP router using http.NewServeMux(). It defines a route for handling POST requests to `/api/students`.

```go

router := http.NewServeMux()
router.HandleFunc("POST /api/students", students.New())

```

## Graceful Shutdown
The server listens for OS signals for shutdown, enabling a smooth shutdown process by allowing any ongoing requests to complete within a timeout of 5 seconds.

```go
done := make(chan os.Signal, 1)
signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

<-done

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
server.Shutdown(ctx)

```
## Logging
The server uses structured logging with slog to log events such as server startup, shutdown, and errors.