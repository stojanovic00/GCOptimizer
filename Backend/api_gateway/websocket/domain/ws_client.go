package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WsClient struct {
	id            uuid.UUID
	socket        *websocket.Conn
	send          chan Message
	Apparatus     Apparatus
	CompetitionId string
}

func (client *WsClient) read(server *WsServer) {
	defer func() {
		server.unregister <- client
		client.socket.Close()
	}()

	//Each received message is broadcast to all other clients
	for {
		var message Message
		err := client.socket.ReadJSON(&message)
		if err != nil {
			break
		}
		jsonMessage, err := json.Marshal(message)
		server.broadcast <- jsonMessage
	}
}

func (client *WsClient) write() {
	defer func() {
		client.socket.Close()
	}()

	for {
		select {
		//ok in this context checks if channel is still open and there is message available
		case message, ok := <-client.send:
			if !ok {
				client.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			client.socket.WriteJSON(message)
		}
	}
}
