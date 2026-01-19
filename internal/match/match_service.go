package match

import (
	"time"
)

type MatchService struct {
}

func (s *MatchService) GetMatches() []Match {
	matches := make([]Match, 0)
	matches = append(matches, Match{
		TeamA:     "Brasil",
		TeamB:     "Argentina",
		MatchType: MatchTypeGroup,
		Round:     1,
		Date:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		Stadium:   "Maracanã",
		// Result:    nil,
		Result: &MatchResult{
			TeamAScore: 1,
			TeamBScore: 1,
		},
	})
	return matches
}
