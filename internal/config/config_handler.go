package config

import (
	"encoding/json"
	"log"
	"net/http"
)

type ConfigHandler struct {
	ConfigService ConfigService // "Dependency Injection"
}

func (h *ConfigHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.GetHealth)
	mux.HandleFunc("POST /health-detailed", h.PostHealthDetailed)
}

// GetHealth godoc
// @Summary      Health Check
// @Description  Get server status
// @Produce      plain
// @Success      200  {string}  string  "Healthy! 2024-01-01 12:00:00"
// @Router       /health [get]
func (h *ConfigHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	healthMsg := h.ConfigService.Health()
	w.Write([]byte(healthMsg))
}

// PostHealthDetailed godoc
// @Summary      Health Check Detailed
// @Description  Get detailed server status
// @Accept       json
// @Produce      json
// @Param        user  body     HealthDetailedRequest  true  "User Data"
// @Success      200  {object}  HealthDetailedResponse
// @Router       /health-detailed [post]
func (h *ConfigHandler) PostHealthDetailed(w http.ResponseWriter, r *http.Request) {
	var req HealthDetailedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("HealthDetailed got %s\n", req.Input)
	healthMsg := h.ConfigService.HealthDetailed()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthMsg)
}
