package utils

import (
	"fmt"
	"net/http"
)

func GenerateAuthKey() string {
	return "Hello"
}

func GetDomainAndKey(r *http.Request) (string, string)  {

	key := r.PathValue("key")
	domain := r.PathValue("domain")
	
	// If there's no domain
	if domain == "" {
		domain = "default"
	}

	return domain, key
}

func CreateRedisKey(domain, key string) string {
	return fmt.Sprintf("%s:%s", domain, key)
}
