package websocket

import (
	scoring_pb "common/proto/scoring/generated"
	"context"
)

type EventHandler struct {
	client scoring_pb.ScoringServiceClient
}

func NewEventHandler(client scoring_pb.ScoringServiceClient) *EventHandler {
	return &EventHandler{client: client}
}

func (h *EventHandler) GetContestantsTempScores(message *EventMessage) *EventResponse {

	response, err := h.client.GetContestantsTempScores(context.Background(), &scoring_pb.ScoreRequest{
		CompetitionId: message.CompetitionId,
		ContestantId:  message.ContestantId,
		Apparatus:     scoring_pb.Apparatus(message.Apparatus),
	})

	if err != nil {
		return &EventResponse{
			Event:         Error,
			Apparatus:     message.Apparatus,
			CompetitionId: message.CompetitionId,
			Response:      &ErrorResponse{Message: err.Error()},
		}
	}

	return &EventResponse{
		Event:         RetrievedContestantsTempScores,
		Apparatus:     message.Apparatus,
		CompetitionId: message.CompetitionId,
		Response:      response.TempScores,
	}
}
func (h *EventHandler) CanCalculateScore(message *EventMessage) *EventResponse {

	response, err := h.client.CanCalculateScore(context.Background(), &scoring_pb.ScoreRequest{
		CompetitionId: message.CompetitionId,
		ContestantId:  message.ContestantId,
		Apparatus:     scoring_pb.Apparatus(message.Apparatus),
	})

	if err != nil {
		return &EventResponse{
			Event:         Error,
			Apparatus:     message.Apparatus,
			CompetitionId: message.CompetitionId,
			Response:      &ErrorResponse{Message: err.Error()},
		}
	}

	return &EventResponse{
		Event:         RetrievedCanCalculate,
		Apparatus:     message.Apparatus,
		CompetitionId: message.CompetitionId,
		Response:      response.IsTrue,
	}
}
func (h *EventHandler) GetScore(message *EventMessage) *EventResponse {

	score, err := h.client.GetScore(context.Background(), &scoring_pb.ScoreRequest{
		CompetitionId: message.CompetitionId,
		ContestantId:  message.ContestantId,
		Apparatus:     scoring_pb.Apparatus(message.Apparatus),
	})

	if err != nil {
		return &EventResponse{
			Event:         Error,
			Apparatus:     message.Apparatus,
			CompetitionId: message.CompetitionId,
			Response:      &ErrorResponse{Message: err.Error()},
		}
	}

	return &EventResponse{
		Event:         RetrievedScore,
		Apparatus:     message.Apparatus,
		CompetitionId: message.CompetitionId,
		Response:      score,
	}
}
func (h *EventHandler) GetNextCurrentApparatusContestant(message *EventMessage) *EventResponse {

	contestant, err := h.client.GetNextCurrentApparatusContestant(context.Background(), &scoring_pb.GetByApparatusRequest{
		CompetitionId: message.CompetitionId,
		Apparatus:     scoring_pb.Apparatus(message.Apparatus),
	})
	if err != nil {
		return &EventResponse{
			Event:         Error,
			Apparatus:     message.Apparatus,
			CompetitionId: message.CompetitionId,
			Response:      &ErrorResponse{Message: err.Error()},
		}
	}

	return &EventResponse{
		Event:         RetrievedNextCurrentApparatusContestant,
		Apparatus:     message.Apparatus,
		CompetitionId: message.CompetitionId,
		Response:      contestant,
	}
}
