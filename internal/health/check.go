package health

import (
	"io"
	"net/http"
)

// HealthCheckHandler  A handler for Health Check API
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{ "status": "pass" }`)
}
