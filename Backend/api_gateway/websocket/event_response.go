package websocket

type EventResponse struct {
	Event         Event       `json:"event,omitempty"`
	Apparatus     Apparatus   `json:"apparatus,omitempty"`
	CompetitionId string      `json:"competitionId,omitempty"`
	ContestantId  string      `json:"contestantId,omitempty"`
	Response      interface{} `json:"response,omitempty"`
}
