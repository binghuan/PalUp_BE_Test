package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const APIKey = "qwerklj1230dsa350123l2k1j4kl1j24"

// Middleware to validate "api-key" in the header
func ValidateAPIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("api-key")
		if apiKey != APIKey {
			http.Error(w, "Unauthorized: Invalid API Key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	// Attach middleware to validate API key
	router.Handle("/test-1", ValidateAPIKeyMiddleware(http.HandlerFunc(Test1))).Methods("POST")

	port := 8082
	fmt.Println("Server is running at port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
