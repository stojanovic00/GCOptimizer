package domain

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type WsServer struct {
	clients    map[*WsClient]bool //bool value is irrelevant, we use map because of convenient search and delete
	broadcast  chan []byte
	register   chan *WsClient
	unregister chan *WsClient
}

func NewWsServer() *WsServer {
	return &WsServer{
		clients:    make(map[*WsClient]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *WsClient),
		unregister: make(chan *WsClient),
	}
}

func (server *WsServer) Start() {
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
			var jsonMessage Message
			json.Unmarshal(message, &jsonMessage)
			server.send(jsonMessage)
		}
	}
}

func (server *WsServer) send(message Message) {
	// Send only to clients with same competitionId and same apparatus
	for client := range server.clients {
		if client.CompetitionId == message.CompetitionId && client.Apparatus == message.Apparatus {
			select {
			case client.send <- message:
			default:
				//If client is unable to receive message we close it its channel and delete it
				close(client.send)
				delete(server.clients, client)
			}
		}
	}
}

func (server *WsServer) OpenConnection(ctx *gin.Context) {
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

	client := &WsClient{
		id:            uuid.New(),
		socket:        connection,
		send:          make(chan Message),
		Apparatus:     Apparatus(apparatus),
		CompetitionId: competitionId,
	}

	server.register <- client

	go client.read(server)
	go client.write()
}
