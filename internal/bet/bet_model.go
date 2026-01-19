package bet

import (
	"bolao-2026/internal/match"
	"fmt"
)

type Bet struct {
	Id      string
	Matches []match.Match
	Score   int
}

func (bet *Bet) CalculateScore(matches []match.Match) {
	score := 0
	errors := make([]error, 0)

	for _, match := range matches {
		// betMatchIdx := slices.IndexFunc(bet.Matches, func(m match.Match) bool { return m.Round == match.Round })
		betMatchIdx := -1
		for i, betMatch := range bet.Matches {
			if betMatch.Round == match.Round {
				betMatchIdx = i
				break
			}
		}
		if betMatchIdx == -1 {
			errors = append(errors, fmt.Errorf("Bet %s: No BetMatch found for round %d", bet.Id, match.Round))
			continue
		}
		betMatch := bet.Matches[betMatchIdx]

		if match.Result == nil {
			continue
		}
		matchWinner := match.Result.Winner()

		if betMatch.Result == nil {
			errors = append(errors, fmt.Errorf("Bet %s: No BetMatch result found for round %d", bet.Id, match.Round))
			continue
		}
		betWinner := betMatch.Result.Winner()

		if betWinner == matchWinner {
			score++
		}
	}

	bet.Score = score
}
