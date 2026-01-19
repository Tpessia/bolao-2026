package bet

import (
	"encoding/json"
	"net/http"
)

type BetHandler struct {
	BetService BetService
}

func (h *BetHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /bets/{id}", h.GetBet)
}

// GetBet godoc
// @Tags         Bet
// @Produce      json
// @Param 			 id path int true "Bet ID"
// @Success      200  {object}  Bet
// @Router       /bets/{id} [get]
func (h *BetHandler) GetBet(w http.ResponseWriter, r *http.Request) {
	betId := r.PathValue("id")
	bet := h.BetService.GetBet(betId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bet)
}
