package main

import (
	"counter/cache"
	"counter/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	defer cache.Rdb.Close()

	http.HandleFunc("/", handler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/count/{domain}/{key}", incrementHandler)
	http.HandleFunc("/count/{key}", incrementHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there, from %s!", r.URL.Path[1:])
}

func createHandler(w http.ResponseWriter, r *http.Request) {

	counter, err := cache.CreateCounter("example2.com", "page_views")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return successfully created counter as a JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(counter)
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	domain, key := utils.GetDomainAndKey(r)
    rKey := utils.CreateRedisKey(domain, key)
    
	fmt.Println("key: ", key, "domain: ", domain)
	fmt.Println("rKey: ", rKey)

	cache.Incr(rKey)
}
