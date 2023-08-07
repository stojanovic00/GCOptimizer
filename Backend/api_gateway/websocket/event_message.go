package websocket

type EventMessage struct {
	Event         Event     `json:"event,omitempty"`
	Apparatus     Apparatus `json:"apparatus,omitempty"`
	CompetitionId string    `json:"competitionId,omitempty"`
	ContestantId  string    `json:"contestantId,omitempty"`
}
