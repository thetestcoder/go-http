# Hints for HTTP Server Implementation

## Understanding the `net/http` Package

Go's `net/http` package provides HTTP client and server implementations. Here are some key components:

### Server Side

- `http.ListenAndServe`: Starts an HTTP server at specified address
- `http.HandleFunc`: Registers a function to handle a specific route
- `http.ResponseWriter`: Interface to construct HTTP responses
- `http.Request`: Represents an HTTP request received by a server

### Client Side

- `http.Get`, `http.Post`, etc.: Functions to make HTTP requests
- `http.Client`: Configurable HTTP client

## Enhancing Your Server

### Adding JSON Responses

```go
func jsonHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{
        "message": "Hello, JSON!",
        "status": "success",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
```

### Serving Static Files

```go
// Serve static files from the "static" directory
fs := http.FileServer(http.Dir("static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

### Request Parameters

```go
func paramHandler(w http.ResponseWriter, r *http.Request) {
    // URL parameters
    query := r.URL.Query()
    name := query.Get("name")

    // POST form data
    r.ParseForm()
    email := r.FormValue("email")

    fmt.Fprintf(w, "Name: %s, Email: %s", name, email)
}
```

### Middleware

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
        next.ServeHTTP(w, r)
    })
}

// Usage
mux := http.NewServeMux()
mux.HandleFunc("/", homeHandler)

loggedMux := loggingMiddleware(mux)
http.ListenAndServe(":8888", loggedMux)
```

## Common Patterns

### Router with Path Variables

For more complex routing, consider using third-party packages like `gorilla/mux` or `httprouter`.

### Graceful Shutdown

```go
server := &http.Server{Addr: ":8888"}

go func() {
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("Server error: %v", err)
    }
}()

// Wait for interrupt signal
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
<-sigChan

// Create a deadline for graceful shutdown
ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
defer cancel()

server.Shutdown(ctx)
log.Println("Server gracefully stopped")
```

## Performance Tips

1. Use response buffering for large responses
2. Implement caching for frequently accessed resources
3. Consider using connection pooling for database connections
4. Set appropriate timeouts for server operations
