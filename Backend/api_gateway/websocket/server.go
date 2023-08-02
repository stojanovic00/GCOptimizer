package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type Server struct {
	clients      map[*Client]bool //bool value is irrelevant, we use map because of convenient search and delete
	broadcast    chan EventMessage
	register     chan *Client
	unregister   chan *Client
	eventHandler *EventHandler
}

func NewServer(eventHandler *EventHandler) *Server {
	return &Server{
		clients:      make(map[*Client]bool),
		broadcast:    make(chan EventMessage),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		eventHandler: eventHandler,
	}
}

func (server *Server) Start() {
	//Runs forever
	for {
		select {
		case client := <-server.register:
			server.clients[client] = true
		case client := <-server.unregister:
			if _, ok := server.clients[client]; ok {
				close(client.send)
				delete(server.clients, client)
			}
		case message := <-server.broadcast:
			response := server.PrepareResponse(&message)
			server.sendToAll(response)
		}
	}
}

func (server *Server) sendToAll(response *EventResponse) {
	// Send only to clients with same competitionId and same apparatus
	for client := range server.clients {
		//If admin sends message everyone gets it, also admin always gets every message
		if (client.CompetitionId == response.CompetitionId && (client.Apparatus == response.Apparatus || client.Apparatus == CompetitionAdmin)) || response.Apparatus == CompetitionAdmin {
			select {
			case client.send <- *response:
			default:
				//If client is unable to receive response we close it its channel and delete it
				close(client.send)
				delete(server.clients, client)
			}
		}
	}
}

func (server *Server) PrepareResponse(message *EventMessage) *EventResponse {
	switch message.Event {
	case TempScoreSubmitted:
		return server.eventHandler.GetContestantsTempScores(message)
	case RetrievedContestantsTempScores:
		return server.eventHandler.CanCalculateScore(message)
	case CalculatedScore, SubmittedScore:
		return server.eventHandler.GetScore(message)
	case ScoredContestant:
		return server.eventHandler.GetNextCurrentApparatusContestant(message)
	case RetrievedNextCurrentApparatusContestant:
		return server.eventHandler.GetContestantsTempScores(message)
	case FinishedRotationOrSession:
		return server.eventHandler.GetCurrentSessionInfo(message)
	case FinishedCompetition:
		return &EventResponse{
			Event:         FinishedCompetition,
			Apparatus:     message.Apparatus,
			CompetitionId: message.CompetitionId,
			Response:      nil,
		}
	default:
		return &EventResponse{
			Event:         Error,
			Apparatus:     message.Apparatus,
			CompetitionId: message.CompetitionId,
			Response:      &ErrorResponse{Message: "Unknown error"},
		}
	}
}

func (server *Server) OpenConnection(ctx *gin.Context) {
	competitionId := ctx.Query("competitionId")
	if competitionId == "" {
		// If "apparatusStr" is not provided in the query, return an error contestant
		ctx.JSON(400, gin.H{"error": "competitionId query parameter is missing"})
		return
	}

	apparatusStr := ctx.Query("apparatus")
	if apparatusStr == "" {
		// If "apparatusStr" is not provided in the query, return an error contestant
		ctx.JSON(400, gin.H{"error": "apparatus query parameter is missing"})
		return
	}
	apparatus, err := strconv.Atoi(apparatusStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid apparatus query parameter"})
		return
	}

	//Upgrade connection to web socket duplex
	//Check origin function is used for CORS (this one allows everything)
	connection, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(ctx.Writer, ctx.Request, nil)
	if error != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	client := &Client{
		id:            uuid.New(),
		socket:        connection,
		send:          make(chan EventResponse),
		Apparatus:     Apparatus(apparatus),
		CompetitionId: competitionId,
	}

	server.register <- client

	go client.read(server)
	go client.write()
}
