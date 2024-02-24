package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()
	proxy := NewReverseProxy("hugo", ":1313")
	os.Setenv("HOST", proxy.host)
	r.Use(proxy.ReverseProxy)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from root"))
	})

	http.ListenAndServe(":8080", r)
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api" {
			w.Write([]byte("Hello from API"))
		} else {
			http.Redirect(w, r, "http://localhost:1313", http.StatusFound)
		}
		next.ServeHTTP(w, r)
	})
}
