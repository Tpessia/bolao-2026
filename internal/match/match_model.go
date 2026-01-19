package match

import "time"

type Match struct {
	TeamA     string
	TeamB     string
	MatchType MatchType
	Round     uint
	Date      time.Time
	Stadium   string
	Result    *MatchResult
}

type MatchResult struct {
	TeamAScore uint
	TeamBScore uint
}

func (mr *MatchResult) Winner() MatchWinner {
	if mr.TeamAScore == mr.TeamBScore {
		return MatchWinnerDraw
	} else if mr.TeamAScore > mr.TeamBScore {
		return MatchWinnerA
	} else {
		return MatchWinnerB
	}
}

type Team struct {
	Name       string
	Group      string
	GroupScore uint
}
