package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func ParseBody(r *http.Request, X interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err = json.Unmarshal(body, X); err != nil {
			return
		}
	}
}

func GetEnvVar(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("db_password is not in .env")
	}
	return value
}
