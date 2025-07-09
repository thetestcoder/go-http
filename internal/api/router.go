package api

import (
	"github.com/thetestcoder/basic-http/internal/middleware"
	"net/http"
)

func setupRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/about", About)
	return middleware.HttpLogging(mux)
}
