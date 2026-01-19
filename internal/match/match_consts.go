package match

type MatchType uint

const (
	MatchTypeGroup MatchType = iota
	MatchTypeKnockout
)

type MatchWinner uint

const (
	MatchWinnerDraw MatchWinner = iota
	MatchWinnerA
	MatchWinnerB
)

type GroupScore uint

const (
	GroupWin  GroupScore = 3
	GroupDraw GroupScore = 1
	GroupLoss GroupScore = 0
)
