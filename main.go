package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		output := fmt.Sprintf("Req: %s %s %s \nData: %v \nFinished", r.Method, r.URL.Host, r.URL.Path, r)
		log.Println(output)
		next.ServeHTTP(w, r)
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("--------------------------------------------")
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	u, _ := url.Parse(os.Getenv("PROXY_URL"))
	http.Handle("/", middlewareOne(middlewareTwo(httputil.NewSingleHostReverseProxy(u))))

	http.ListenAndServe(":"+port, nil)
}
