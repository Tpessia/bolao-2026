package bet

import match "bolao-2026/internal/match"

type BetService struct {
	MatchService match.MatchService
}

func (bs *BetService) GetBet(id string) Bet {
	betMatches := bs.MatchService.GetMatches()

	matches := bs.MatchService.GetMatches()
	matches[0].Result = &match.MatchResult{
		TeamAScore: 2,
		TeamBScore: 1,
	}

	bet := Bet{
		Id:      "test",
		Matches: matches,
	}

	bet.CalculateScore(betMatches)

	return bet
}
