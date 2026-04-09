package routes

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"api-gateway/middleware"
)

func proxyRequest(target string) http.Handler {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
}

func RegisterRoutes(mux *http.ServeMux) {

	// Public routes
	mux.Handle("/auth/", proxyRequest("http://auth-service:3000"))

	// Example protected route
	mux.Handle("/protected/",
		middleware.JWTAuth(proxyRequest("http://auth-service:3000")),
	)
}