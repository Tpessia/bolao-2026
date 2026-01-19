package match

import (
	"encoding/json"
	"net/http"
)

type MatchHandler struct {
	MatchService MatchService
}

func (h *MatchHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /matches", h.GetMatches)
}

// GetMatches godoc
// @Tags         Match
// @Produce      json
// @Success      200  {object}  []Match
// @Router       /matches [get]
func (h *MatchHandler) GetMatches(w http.ResponseWriter, r *http.Request) {
	matches := h.MatchService.GetMatches()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}
