package handlers

import (
	"encoding/json"
	"hexagonal-example/internal/core/ports"
	"net/http"
)

type HTTPHealthHandler struct {
	HealthService ports.HealthInterface
}

func (h HTTPHealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status := h.HealthService.IsAppHealthy()

	json.NewEncoder(w).Encode(status)
}