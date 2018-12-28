package utils

import (
	"net/http"

	"github.com/spf13/viper"
)

// GetClient - returns the default HTTP Client
func GetClient() *http.Client {
	client := &http.Client{}
	return client
}

// SetDefaultHeaders - adds the default HTTP Headers to the Request
func SetDefaultHeaders(req *http.Request) {
	req.Header.Add("Private-Token", viper.GetString("Token"))
	req.Header.Add("accept", "application/json, text/plain, */*")
}
