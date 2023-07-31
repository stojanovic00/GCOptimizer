package websocket

type Event int8

const (
	Error Event = iota
	TempScoreSubmitted
	RetrievedContestantsTempScores
)
