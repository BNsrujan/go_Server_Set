package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	// we are taking the http.ResponseWriter and writing w to it
	// fmt formate
	// net/http to create a server

	// if you are not using gin you have write this many qutions
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		re := regexp.MustCompile(`^/user/(\d+)$`)
		matches := re.FindStringSubmatch(r.URL.Path)

		if len(matches) < 2 {
			http.Error(w, "Invalid user ID", http.StatusBadGateway)
		}

		id := matches[1]

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"user_id": id})
	})

	server.ListenAndServe()
}
