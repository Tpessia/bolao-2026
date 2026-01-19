package bet

import (
	match "bolao-2026/internal/match"
	"testing"
	"time"
)

func TestBetWin(t *testing.T) {
	matches := getMatches()

	betMatches := getMatches()

	bet := Bet{
		Id:      "test",
		Matches: betMatches,
	}

	bet.CalculateScore(matches)

	passes := bet.Score == 3
	if !passes {
		t.Errorf("Bet1 failed with %d", bet.Score)
	}
}

func TestBetLoss(t *testing.T) {
	matches := getMatches()

	betMatches := getMatches()
	betMatches[1].Result = &match.MatchResult{
		TeamAScore: 1,
		TeamBScore: 1,
	}

	bet := Bet{
		Id:      "test",
		Matches: betMatches,
	}

	bet.CalculateScore(matches)

	passes := bet.Score == 2
	if !passes {
		t.Errorf("Bet1 failed with %d", bet.Score)
	}
}

func getMatches() []match.Match {
	matches := make([]match.Match, 0)
	matches = append(matches, match.Match{
		TeamA:     "Brasil",
		TeamB:     "Argentina",
		MatchType: match.MatchTypeGroup,
		Round:     1,
		Date:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		Stadium:   "Maracanã",
		Result: &match.MatchResult{
			TeamAScore: 1,
			TeamBScore: 1,
		},
	})
	matches = append(matches, match.Match{
		TeamA:     "Brasil",
		TeamB:     "Holanda",
		MatchType: match.MatchTypeGroup,
		Round:     2,
		Date:      time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
		Stadium:   "Maracanã",
		Result: &match.MatchResult{
			TeamAScore: 2,
			TeamBScore: 1,
		},
	})
	matches = append(matches, match.Match{
		TeamA:     "Brasil",
		TeamB:     "Peru",
		MatchType: match.MatchTypeGroup,
		Round:     3,
		Date:      time.Date(2026, 1, 3, 0, 0, 0, 0, time.UTC),
		Stadium:   "Maracanã",
		Result: &match.MatchResult{
			TeamAScore: 1,
			TeamBScore: 2,
		},
	})
	matches = append(matches, match.Match{
		TeamA:     "Brasil",
		TeamB:     "França",
		MatchType: match.MatchTypeGroup,
		Round:     4,
		Date:      time.Date(2026, 1, 4, 0, 0, 0, 0, time.UTC),
		Stadium:   "Maracanã",
		Result:    nil,
	})
	return matches
}
